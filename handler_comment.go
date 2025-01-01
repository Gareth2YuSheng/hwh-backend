package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// UNVOTE COMMENT
func (apiCfg *APIConfig) handlerUnVoteComment(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerUnVoteComment")
	//Get CommentID from URL Params
	commentIdStr := chi.URLParam(r, "commentID")
	commentId, err := uuid.Parse(commentIdStr)
	if err != nil {
		logError("Unable to Get CommentID", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Vote Comment: Invalid CommentId")
		return
	}

	fmt.Printf("Comment: %v\n", commentId) //remove later
	// fmt.Printf("Body: %v\n", req)          //remove later

	//Check whether vote already exists
	vote, err := apiCfg.DB.GetVotesForCommentByUser(commentId, user.UserID)
	if err != nil {
		logError("Error checking if vote exists", err)
		respondERROR(w, http.StatusBadRequest, "Failed to UnVote Comment: User has not voted on this comment before")
		return
	}

	//Delete Comment
	err = apiCfg.DB.DeleteVote(vote)
	if err != nil {
		logError("Unable to Delete Vote", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to UnVote Comment")
		return
	}

	oppositeVoteValue := 0
	if vote.VoteValue == 1 {
		oppositeVoteValue = -1
	} else if vote.VoteValue == -1 {
		oppositeVoteValue = 1
	}

	//Update Comment Vote Count
	err = apiCfg.DB.UpdateCommentVoteCountByCommentID(commentId, oppositeVoteValue)
	if err != nil {
		logError(fmt.Sprintf("UnvoteComment Handler - Unable to Update Comment [%d] Vote Count", commentId), err)
		respondERROR(w, http.StatusInternalServerError, "Failed to UnVote Comment")
		return
	}

	respondOK(w, http.StatusOK, "Comment Unvoted Successfully", nil)
}

// VOTE COMMENT
func (apiCfg *APIConfig) handlerVoteComment(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerVoteComment")
	req := VoteCommentRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorParsingJSON(err, w)
		return
	}
	defer r.Body.Close()

	//Get CommentID from URL Params
	commentIdStr := chi.URLParam(r, "commentID")
	commentId, err := uuid.Parse(commentIdStr)
	if err != nil {
		logError("Unable to Get CommentID", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Vote Comment: Invalid CommentId")
		return
	}

	fmt.Printf("Comment: %v\n", commentId) //remove later
	fmt.Printf("Body: %v\n", req)          //remove later

	//Validation
	if strings.ToLower(req.VoteType) != "up" && strings.ToLower(req.VoteType) != "down" {
		logError("Invalid Vote Type", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Vote Comment: Invalid Vote Type")
		return
	}

	//Check whether vote already exists
	vote, err := apiCfg.DB.GetVotesForCommentByUser(commentId, user.UserID)
	if err != nil {
		if !strings.HasSuffix(err.Error(), "not found") {
			logError("Error checking if vote exists", err)
			respondERROR(w, http.StatusBadRequest, "Failed to Vote Comment")
			return
		}
		// if not found error just continue because it is expected behavior
	}

	voteVal := 0
	if strings.ToLower(req.VoteType) == "up" {
		voteVal = 1
	} else if strings.ToLower(req.VoteType) == "down" {
		voteVal = -1
	}

	fmt.Printf("Vote Value: %d\n", voteVal) //remove later

	//If vote already exists and has the same vote value - do nothing
	if vote != nil && vote.VoteValue == voteVal {
		respondOK(w, http.StatusOK, "Comment Voted Successfully", nil)
		return
	}
	//Else if vote already exists but has a different vote value - update vote
	if vote != nil && vote.VoteValue != voteVal {
		//Update Vote Value
		err = vote.UpdateVoteValue(voteVal)
		if err != nil {
			logError(fmt.Sprintf("Unable to Update Vote [%d]", vote.VoteID), err)
			respondERROR(w, http.StatusBadRequest, "Failed to Update Vote Value: Invalid Vote Details")
			return
		}

		//Update Vote
		err = apiCfg.DB.UpdateVoteVoteValue(vote)
		if err != nil {
			logError(fmt.Sprintf("Unable to Update Vote [%d] Value", commentId), err)
			respondERROR(w, http.StatusInternalServerError, "Failed to Update Vote Value")
			return
		}

		//Update vote value for updating comment vote count
		voteVal *= 2
	} else {
		//Else if vote does not exist - create vote
		//Check Comment Exists First
		comment, err := apiCfg.DB.GetCommentByCommentID(commentId)
		if err != nil {
			logError(fmt.Sprintf("Unable to Get Comment [%d]", commentId), err)
			respondERROR(w, http.StatusBadRequest, "Failed to Update Comment: Invalid CommentId")
			return
		}

		vote, err = NewVote(comment.CommentID, user.UserID, voteVal)
		if err != nil {
			logError("Error Creating New Vote Template", err)
			respondERROR(w, http.StatusBadRequest, "Failed to Create Vote, Invalid Vote Details")
			return
		}

		//Create Vote
		err = apiCfg.DB.CreateVote(vote)
		if err != nil {
			logError("Unable to Create Vote", err)
			respondERROR(w, http.StatusBadRequest, "Failed to Create Vote")
			return
		}
	}

	fmt.Printf("Vote Value Update by: %d\n", voteVal) //remove later

	//Update Comment Vote Count for last 2 cases
	err = apiCfg.DB.UpdateCommentVoteCountByCommentID(commentId, voteVal)
	if err != nil {
		logError(fmt.Sprintf("Unable to Update Comment [%d] Vote Count", commentId), err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Update Comment")
		return
	}

	respondOK(w, http.StatusOK, "Comment Voted Successfully", nil)
}

// CREATE COMMENT
func (apiCfg *APIConfig) handlerCreateComment(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerCreateComment")
	req := CreateCommentRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorParsingJSON(err, w)
		return
	}
	defer r.Body.Close()

	//Get ThreadID from URL Params
	threadIDStr := chi.URLParam(r, "threadID")
	threadID, err := uuid.Parse(threadIDStr)
	if err != nil {
		logError("Unable to Get ThreadID", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Create Comment: Invalid ThreadId")
		return
	}

	//Check whether thread exists
	_, err = apiCfg.DB.GetThreadByThreadID(threadID)
	if err != nil {
		logError(fmt.Sprintf("Thread [%v] does not exist", threadID), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Create Comment, Invalid ThreadId")
		return
	}

	comment, err := NewComment(req.Content, threadID, user.UserID)
	if err != nil {
		logError("Error Creating New Comment Template", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Create Comment, Invalid Comment Details")
		return
	}
	fmt.Printf("New Comment: %v\n", comment) //remove later

	//Create Comment
	err = apiCfg.DB.CreateComment(comment)
	if err != nil {
		logError("Unable to Create Comment", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Create Comment")
		return
	}

	//Update Thead Comment Count
	err = apiCfg.DB.UpdateThreadCommentCountByThreadID(threadID, 1)
	if err != nil {
		logError(fmt.Sprintf("Unable to Update Thead [%v] Comment Count", threadID), err)
		respondERROR(w, http.StatusInternalServerError, "Something Went Wrong")
		return
	}

	respondOK(w, http.StatusCreated, "Comment Created Successfully", nil)
}

// GET ALL COMMENTS FOR THREAD
func (apiCfg *APIConfig) handlerGetAllComments(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerGetAllComments")
	//Get Compulsory URL Query Params
	fmt.Printf("%v\n", r.URL.Query()) //remove later
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		InvalidURLQuery("GetAllComments: Unable to get Count from URL Query", err, w)
		return
	}
	if count == 0 {
		InvalidURLQuery("GetAllComments: Count cannot be 0", err, w)
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

	//Get ThreadID from URL Params
	threadIDStr := chi.URLParam(r, "threadID")
	threadID, err := uuid.Parse(threadIDStr)
	if err != nil {
		logError("Unable to Get ThreadID", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Create Comment: Invalid ThreadId")
		return
	}

	//Check whether thread exists
	_, err = apiCfg.DB.GetThreadByThreadID(threadID)
	if err != nil {
		logError(fmt.Sprintf("Thread [%v] does not exist", threadID), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Create Comment, Invalid ThreadId")
		return
	}

	comments, commentCount, err := apiCfg.DB.GetAllCommentsByThreadIDWithVotesByUserID(count, page, threadID, user.UserID)
	// comments, commentCount, err := apiCfg.DB.GetAllCommentsByThreadID(count, page, threadID)
	if err != nil {
		logError(fmt.Sprintf("Unable to Get Comments: Page[%d] Count[%d] ThreadID[%v]", page, count, threadID), err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Get Comments")
		return
	}

	respondOK(w, http.StatusOK, "", GetCommentsWithVoteResponse{
		CommentCount: commentCount,
		Comments:     comments,
	})
}

// UPDATE COMMENT
func (apiCfg *APIConfig) handlerUpdateComment(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerUpdateComment")
	req := CreateCommentRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorParsingJSON(err, w)
		return
	}
	defer r.Body.Close()

	//Get CommentID from URL Params
	commentIdStr := chi.URLParam(r, "commentID")
	commentId, err := uuid.Parse(commentIdStr)
	if err != nil {
		logError("Unable to Get CommentID", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Update Comment: Invalid CommentId")
		return
	}

	fmt.Printf("Comment: %v\n", commentId) //remove later
	fmt.Printf("Body: %v\n", req)          //remove later

	//Check Comment Exists First
	comment, err := apiCfg.DB.GetCommentByCommentID(commentId)
	if err != nil {
		logError(fmt.Sprintf("Unable to Get Comment [%d]", commentId), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Update Comment: Invalid CommentId")
		return
	}

	//Validate User if they are the author of the comment
	if comment.AuthorID != user.UserID {
		logError(fmt.Sprintf("User [%d] does not have permission to edit Comment [%d]", user.UserID, commentId), err)
		PermissionDeniedRes(w)
		return
	}

	//Update Comment Details
	err = comment.UpdateCommentContent(req.Content)
	if err != nil {
		logError(fmt.Sprintf("Unable to Update Comment [%d]", commentId), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Update Comment: Invalid Comment Details")
		return
	}

	err = apiCfg.DB.UpdateCommentContent(comment)
	if err != nil {
		logError(fmt.Sprintf("Unable to Update Comment [%d]", commentId), err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Update Comment")
		return
	}

	respondOK(w, http.StatusOK, "Comment Updated Successfully", nil)
}

// MARK COMMENT AS CORRECT
func (apiCfg *APIConfig) handlerMarkCommentAsAnswer(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerMarkCommentAsAnswer")
	//User Req Query instead of Req Body as Boolean will always default to false if missing in body
	fmt.Printf("%v\n", r.URL.Query()) //remove later
	isAnswerStr := strings.ToLower(r.URL.Query().Get("isAnswer"))
	//Validation
	if isAnswerStr != "true" && isAnswerStr != "false" {
		InvalidURLQuery("MarkCommentAsAnswer: Unable to get isAnswer from URL Query", nil, w)
		return
	}
	isAnswer, err := strconv.ParseBool(isAnswerStr)
	if err != nil {
		InvalidURLQuery("MarkCommentAsAnswer: Unable to get isAnswer from URL Query", err, w)
		return
	}

	//Get CommentID from URL Params
	commentIdStr := chi.URLParam(r, "commentID")
	commentId, err := uuid.Parse(commentIdStr)
	if err != nil {
		logError("Unable to Get CommentID", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Mark Comment as Answer: Invalid CommentId")
		return
	}

	fmt.Printf("Comment: %v\n", commentId) //remove later
	// fmt.Printf("Body: %v\n", req)          //remove later

	//Check Comment Exists First
	comment, err := apiCfg.DB.GetCommentByCommentID(commentId)
	if err != nil {
		logError(fmt.Sprintf("Unable to Get Comment [%d]", commentId), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Mark Comment as Answer: Invalid CommentId")
		return
	}

	//Get Thread AuthorID
	thread, err := apiCfg.DB.GetThreadByThreadID(comment.ThreadID)
	if err != nil {
		logError(fmt.Sprintf("Unable to Get Thread [%d] to Mark Comment as Correct", comment.ThreadID), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Mark Comment as Answer: Invalid CommentId")
		return
	}

	//Validate User if they are the author of the thread !!!!
	if thread.AuthorID != user.UserID {
		logError(fmt.Sprintf("User [%d] does not have permission to Mark Comment [%d] as Correct", user.UserID, commentId), err)
		PermissionDeniedRes(w)
		return
	}

	//Check if new bool is same as existing
	if comment.IsAnswer == isAnswer {
		respondOK(w, http.StatusOK, "Comment Marked as Answer Successfully", nil)
		return
	}

	//Update Comment Details
	err = comment.UpdateCommentIsAnswer(isAnswer)
	if err != nil {
		logError(fmt.Sprintf("Unable to Mark Comment [%d] as Answer", commentId), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Mark Comment as Answer: Invalid Comment Details")
		return
	}

	err = apiCfg.DB.UpdateCommentIsAnswer(comment)
	if err != nil {
		logError("Unable to Update Comment", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Mark Comment as Answer")
		return
	}

	respondOK(w, http.StatusOK, "Comment Marked as Answer Successfully", nil)
}

// DELETE COMMENT
func (apiCfg *APIConfig) handlerDeleteComment(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerDeleteComment")
	commentIdStr := chi.URLParam(r, "commentID")
	commentId, err := uuid.Parse(commentIdStr)
	if err != nil {
		logError("Unable to Get CommentID", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Update Comment: Invalid CommentId")
		return
	}

	fmt.Printf("Comment: %v\n", commentId) //remove later

	//Check Comment Exists First
	comment, err := apiCfg.DB.GetCommentByCommentID(commentId)
	if err != nil {
		logError(fmt.Sprintf("Unable to Get Comment [%d]", commentId), err)
		respondERROR(w, http.StatusBadRequest, "Failed to Delete Comment: Invalid CommentId")
		return
	}

	//Validate User if they are the author of the thread OR an ADMIN
	if comment.AuthorID != user.UserID && user.Role != "Admin" {
		logError(fmt.Sprintf("User [%d] does not have permission to delete Comment [%d]", user.UserID, commentId), err)
		PermissionDeniedRes(w)
		return
	}

	//Delete Comment
	err = apiCfg.DB.DeleteCommentByCommentID(comment.CommentID)
	if err != nil {
		logError("Unable to Delete Comment", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Delete Comment")
		return
	}

	//Update Thead Comment Count
	err = apiCfg.DB.UpdateThreadCommentCountByThreadID(comment.ThreadID, -1)
	if err != nil {
		logError(fmt.Sprintf("Unable to Update Thead [%v] Comment Count", comment.ThreadID), err)
		respondERROR(w, http.StatusInternalServerError, "Something Went Wrong")
		return
	}

	respondOK(w, http.StatusOK, "Comment Deleted Successfully", nil)
}
