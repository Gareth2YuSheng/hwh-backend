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

	err = apiCfg.DB.CreateThread(thread)
	if err != nil {
		logError("Unable to Create Thread", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Create Thread")
		return
	}

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

	threadIDStr := chi.URLParam(r, "threadID")
	threadID, err := uuid.Parse(threadIDStr)
	if err != nil {
		logError("Unable to Get ThreadID", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Update Thread: Invalid ThreadId")
		return
	}

	fmt.Printf("Thread: %v\n", threadID) //remove later
	fmt.Printf("Body: %v\n", req)        //remove later

	//Check Thread Exists First
	thread, err := apiCfg.DB.GetThreadByThreadID(threadID)
	if err != nil {
		logError(fmt.Sprintf("Unable to Get Thread [%d]", threadID), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Update Thread: Invalid ThreadId")
		return
	}

	//Update Thread Details
	thread.UpdateThread(req.Title, req.Content)

	err = apiCfg.DB.UpdateThread(thread)
	if err != nil {
		logError("Unable to Update Thread", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Update Thread")
		return
	}

	respondOK(w, http.StatusCreated, "Thread Updated Successfully", nil)
}

// UPDATE THREAD
func (apiCfg *APIConfig) handlerDeleteThread(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerDeleteThread")
	threadIDStr := chi.URLParam(r, "threadID")
	threadID, err := uuid.Parse(threadIDStr)
	if err != nil {
		logError("Unable to Get ThreadID", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Delete Thread: Invalid ThreadId")
		return
	}

	fmt.Printf("Thread: %v\n", threadID) //remove later

	//Check Thread Exists First
	thread, err := apiCfg.DB.GetThreadByThreadID(threadID)
	if err != nil {
		logError(fmt.Sprintf("Unable to Get Thread [%d]", threadID), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Delete Thread: Invalid ThreadId")
		return
	}

	err = apiCfg.DB.DeleteThreadByThreadID(thread.ThreadID)
	if err != nil {
		logError("Unable to Delete Thread", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Delete Thread")
		return
	}

	respondOK(w, http.StatusCreated, "Thread Deleted Successfully", nil)
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

	threads, err := apiCfg.DB.GetAllThreads(count, page, search, tagId)
	if err != nil {
		logError(fmt.Sprintf("Unable to Get Threads: Page[%d] Count[%d] Search[%v] TagID[%v]", page, count, search, tagId), err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Get Threads")
		return
	}

	respondOK(w, http.StatusOK, "", GetThreadsResponse{
		Threads: threads,
	})
	// respondOK(w, http.StatusOK, "", struct{}{})
}

// GET DETAILS OF A SPECIFIC THREAD WITH COMMENTS
func (apiCfg *APIConfig) handlerGetTheadDetails(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerGetTheadDetails")
	//Get Compulsory URL Query Params
	// fmt.Printf("%v\n", r.URL.Query()) //remove later
	// count, err := strconv.Atoi(r.URL.Query().Get("count"))
	// if err != nil {
	// 	InvalidURLQuery("GetAllThreads: Unable to get Count from URL Query", err, w)
	// 	return
	// }
	// if count == 0 {
	// 	InvalidURLQuery("GetAllThreads: Count cannot be 0", err, w)
	// 	return
	// }
	// fmt.Println("Count: ", count) //remove later
	// page, err := strconv.Atoi(r.URL.Query().Get("page"))
	// if err != nil {
	// 	InvalidURLQuery("GetAllThreads: Unable to get Page from URL Query", err, w)
	// 	return
	// }
	// if page == 0 {
	// 	InvalidURLQuery("GetAllThreads: Page cannot be 0", err, w)
	// 	return
	// }
	// fmt.Println("Page: ", page) //remove later

	// //Check and Get Optional URL Query Params
	// search := ""
	// _, searchOk := r.URL.Query()["search"]
	// if searchOk {
	// 	search = r.URL.Query().Get("search")
	// 	fmt.Println("Search Title: ", search) //remove later
	// }
	// tagId := uuid.Nil
	// _, tagOK := r.URL.Query()["tagId"]
	// if tagOK {
	// 	tagId, err = uuid.Parse(r.URL.Query().Get("tagId"))
	// 	if err != nil {
	// 		InvalidURLQuery("GetAllThreads: Unable to get TagID from URL Query", err, w)
	// 		return
	// 	}
	// 	fmt.Println("Filter By TagID: ", tagId) //remove later
	// }
	// if search == "" {
	// 	fmt.Println("No Search Title: ", search) //remove later
	// }
	// if tagId == uuid.Nil {
	// 	fmt.Println("No Filter By TagID: ", tagId) //remove later
	// }

	// threads, err := apiCfg.DB.GetAllThreads(count, page, search, tagId)
	// if err != nil {
	// 	logError(fmt.Sprintf("Unable to Get Threads: Page[%d] Count[%d] Search[%v] TagID[%v]", page, count, search, tagId), err)
	// 	respondERROR(w, http.StatusInternalServerError, "Failed to Get Threads")
	// 	return
	// }

	// respondOK(w, http.StatusOK, "", GetThreadsResponse{
	// 	Threads: threads,
	// })
	respondOK(w, http.StatusOK, "", struct{}{})
}
