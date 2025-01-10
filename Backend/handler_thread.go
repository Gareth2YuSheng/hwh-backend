package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// CREATE NEW THREAD
func (apiCfg *APIConfig) handlerCreateThread(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerCreateThread")
	req := CreateThreadRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorParsingJSON(err, w)
		return
	}
	defer r.Body.Close()

	thread, err := NewThread(req.Title, req.Content, user.UserID, req.TagID)
	if err != nil {
		logError("Error Creating New Standard Thread Template", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Create Thread, Invalid Thread Details")
		return
	}
	fmt.Printf("New Thread: %v\n", thread) //remove later

	//Create Thread
	err = apiCfg.DB.CreateThread(thread)
	if err != nil {
		logError("Unable to Create Thread", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Create Thread")
		return
	}

	//Update Total Thread Tally Count
	// err = apiCfg.DB.UpdateTotalThreadTally(1)
	// if err != nil {
	// 	logError("Unable to Update Total Thead Tally", err)
	// 	respondERROR(w, http.StatusInternalServerError, "Something Went Wrong")
	// 	return
	// }

	// //Update Thread Tally Count for TagID
	// err = apiCfg.DB.UpdateTagThreadTally(req.TagID, 1)
	// if err != nil {
	// 	logError(fmt.Sprintf("Unable to Update Thead Tally for Tag [%v]", req.TagID), err)
	// 	respondERROR(w, http.StatusInternalServerError, "Something Went Wrong")
	// 	return
	// }

	respondOK(w, http.StatusCreated, "Thread Created Successfully", nil)
}

// UPDATE THREAD
func (apiCfg *APIConfig) handlerUpdateThread(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerUpdateThread")
	req := UpdateThreadRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorParsingJSON(err, w)
		return
	}
	defer r.Body.Close()

	//Get ThreadID from URL Params
	threadIdStr := chi.URLParam(r, "threadID")
	threadId, err := uuid.Parse(threadIdStr)
	if err != nil {
		logError("Unable to Get ThreadID", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Update Thread: Invalid ThreadId")
		return
	}

	fmt.Printf("Thread: %v\n", threadId) //remove later
	fmt.Printf("Body: %v\n", req)        //remove later

	//Check Thread Exists First
	thread, err := apiCfg.DB.GetThreadByThreadID(threadId)
	if err != nil {
		logError(fmt.Sprintf("Unable to Get Thread [%s]", threadId.String()), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Update Thread: Invalid ThreadId")
		return
	}

	//Validate User if they are the author of the thread
	if thread.AuthorID != user.UserID {
		logError(fmt.Sprintf("User [%s] does not have permission to edit Thread [%s]", user.UserID.String(), threadId.String()), err)
		PermissionDeniedRes(w)
		return
	}

	//Update Thread Details
	err = thread.UpdateThread(req.Title, req.Content)
	if err != nil {
		logError(fmt.Sprintf("Unable to Update Thread [%s]", threadId.String()), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Update Thread: Invalid Thread Details")
		return
	}

	err = apiCfg.DB.UpdateThread(thread)
	if err != nil {
		logError("Unable to Update Thread", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Update Thread")
		return
	}

	respondOK(w, http.StatusOK, "Thread Updated Successfully", nil)
}

// DELETE THREAD
func (apiCfg *APIConfig) handlerDeleteThread(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerDeleteThread")
	threadIdStr := chi.URLParam(r, "threadID")
	threadId, err := uuid.Parse(threadIdStr)
	if err != nil {
		logError("Unable to Get ThreadID", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Delete Thread: Invalid ThreadId")
		return
	}

	fmt.Printf("Thread: %v\n", threadId) //remove later

	//Check Thread Exists First
	thread, err := apiCfg.DB.GetThreadByThreadID(threadId)
	if err != nil {
		logError(fmt.Sprintf("Unable to Get Thread [%s]", threadId.String()), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Delete Thread: Invalid ThreadId")
		return
	}

	//Validate User if they are the author of the thread OR an ADMIN
	if thread.AuthorID != user.UserID && user.Role != "Admin" {
		logError(fmt.Sprintf("User [%s] does not have permission to delete Thread [%s]", user.UserID.String(), threadId.String()), err)
		PermissionDeniedRes(w)
		return
	}

	err = apiCfg.DB.DeleteThreadByThreadID(thread.ThreadID)
	if err != nil {
		logError("Unable to Delete Thread", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Delete Thread")
		return
	}

	respondOK(w, http.StatusOK, "Thread Deleted Successfully", nil)
}

// GET ALL THREADS
func (apiCfg *APIConfig) handlerGetAllThreads(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerGetAllThreads")
	//Get Compulsory URL Query Params
	fmt.Printf("%v\n", r.URL.Query()) //remove later
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		InvalidURLQuery("GetAllThreads: Unable to get Count from URL Query", err, w)
		return
	}
	if count == 0 {
		InvalidURLQuery("GetAllThreads: Count cannot be 0", err, w)
		return
	}
	fmt.Println("Count: ", count) //remove later
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		InvalidURLQuery("GetAllThreads: Unable to get Page from URL Query", err, w)
		return
	}
	if page == 0 {
		InvalidURLQuery("GetAllThreads: Page cannot be 0", err, w)
		return
	}
	fmt.Println("Page: ", page) //remove later

	//Check and Get Optional URL Query Params
	search := ""
	_, searchOk := r.URL.Query()["search"]
	if searchOk {
		search = r.URL.Query().Get("search")
		fmt.Println("Search Title: ", search) //remove later
	}
	tagId := uuid.Nil
	_, tagOK := r.URL.Query()["tagId"]
	if tagOK {
		tagId, err = uuid.Parse(r.URL.Query().Get("tagId"))
		if err != nil {
			InvalidURLQuery("GetAllThreads: Unable to get TagID from URL Query", err, w)
			return
		}
		fmt.Println("Filter By TagID: ", tagId) //remove later
	}
	if search == "" {
		fmt.Println("No Search Title: ", search) //remove later
	}
	if tagId == uuid.Nil {
		fmt.Println("No Filter By TagID: ", tagId) //remove later
	}

	threads, threadCount, err := apiCfg.DB.GetAllThreads(count, page, search, tagId)
	if err != nil {
		logError(fmt.Sprintf("Unable to Get Threads: Page[%d] Count[%d] Search[%v] TagID[%v]", page, count, search, tagId), err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Get Threads")
		return
	}

	// tally, err := apiCfg.DB.GetTotalThreadTally()
	// if err != nil {
	// 	logError("Unable to Get Total Thread Tally", err)
	// 	respondERROR(w, http.StatusInternalServerError, "Failed to Get Threads")
	// 	return
	// }

	respondOK(w, http.StatusOK, "", GetThreadsResponse{
		ThreadCount: threadCount,
		Threads:     threads,
	})
	// respondOK(w, http.StatusOK, "", struct{}{})
}

// GET DETAILS OF A SPECIFIC THREAD
func (apiCfg *APIConfig) handlerGetTheadDetails(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerGetTheadDetails")
	//Get ThreadID from URL Params
	threadIDStr := chi.URLParam(r, "threadID")
	threadID, err := uuid.Parse(threadIDStr)
	if err != nil {
		logError("Unable to Get ThreadID", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Get Thread: Invalid ThreadId")
		return
	}

	thread, err := apiCfg.DB.GetThreadDetailsByThreadID(threadID)
	if err != nil {
		logError(fmt.Sprintf("Unable to Get Thread [%s]", threadID.String()), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Get Thread: Invalid ThreadId")
		return
	}

	respondOK(w, http.StatusOK, "", GetThreadDetailsResponse{
		Thread: thread,
	})
}
