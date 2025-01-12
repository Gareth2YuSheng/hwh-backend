import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface Vote {
  voteId: string;
  voteValue: number;
}

export interface Comment {
  commentId: string;
  content: string;
  voteCount: number;
  authorId: string;
  author: string;
  threadId: string;
  createdAt: string;
  updatedAt: string;
  isAnswer: boolean;
  vote: Vote;
}

interface CommentsState {
  comments: Comment[];
  totalComments: number;
  isLoading: boolean;
  error: string | null;
  comment: Comment | null;
}

const initialState: CommentsState = {
  comments: [],
  totalComments: 0,
  isLoading: false,
  error: null,
  comment: null
}

export const fetchCommentData = createAsyncThunk("thread/fetchCommentData", async ({ token, page, count, threadId }: 
  { token: string | undefined, page: number, count: number, threadId: string | undefined }) => {
  try {
    const response = await fetch(`http://localhost:8080/comment/${threadId}?count=${count}&page=${page}`, {
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${token}`
      }
    });
    const content = await response.json();
    console.log(content.data)
    return content.data;
  } catch (err) {
    console.log("Error:", err);
  }
});

export const commentSlice = createSlice({
  name: "comment",
  initialState: initialState,
  reducers: {
    selectComment: (state, action:PayloadAction<{ index: number; }>) => {
      state.comment = state.comments[action.payload.index];
    },
    voteComment: (state, action:PayloadAction<{ index: number; amt: number; }>) => {
      state.comments[action.payload.index].voteCount += action.payload.amt;
    },
    markAnswerComment: (state, action:PayloadAction<{ index: number; isAns: boolean; }>) => {
      state.comments[action.payload.index].isAnswer = action.payload.isAns;
    }
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchCommentData.pending, (state) => {
        console.log("Getting comment data pending");
        state.isLoading = true;
      })
      .addCase(fetchCommentData.fulfilled, (state, action) => {
        console.log("Getting comment data fulfilled");
        state.error = null;
        state.isLoading = false;
        state.comments = action.payload.comments;
        state.totalComments = action.payload.commentCount;
      })
      .addCase(fetchCommentData.rejected, (state, action) => {
        console.log("Getting comment data rejected");
        state.isLoading = false;
        state.error = action.error.message || "Failed to fetch Comments";
      });
  }
});

export const { selectComment, voteComment, markAnswerComment } = commentSlice.actions;

export default commentSlice.reducer;