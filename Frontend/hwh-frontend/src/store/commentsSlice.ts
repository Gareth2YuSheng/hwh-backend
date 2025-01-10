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
}

const initialState: CommentsState = {
  comments: [],
  totalComments: 0,
  isLoading: false,
  error: null
}

export const fetchCommentData = createAsyncThunk("thread/fetchCommentData", async ({ token, page, count, tagId, search }: 
  { token: string | undefined, page: number, count: number, tagId: string, search: string }) => {
  let url = `http://localhost:8080/thread/all?count=${count}&page=${page}`;
  if (tagId != "") {
    url += `&tagId=${tagId}`;
  }
  if (search != "") {
    url += `&search=${search}`;
  }
  try {
    const response = await fetch(url, {
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${token}`
      }
    });
    const content = await response.json();
    return content.data;
  } catch (err) {
    console.log("Error:", err);
  }
});

export const commentSlice = createSlice({
  name: "comment",
  initialState: initialState,
  reducers: {
    // addHabit: (state, action:PayloadAction<{name:string; frequency:"daily"|"weekly"}>) => {
    //   const newHabit: Habit = {
    //     id: Date.now().toString(),
    //     name: action.payload.name,
    //     frequency: action.payload.frequency,
    //     completedDates: [],
    //     createdAt: new Date().toISOString(),
    //   }
    //   state.habits.push(newHabit);
    // },
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

export const {  } = commentSlice.actions;

export default commentSlice.reducer;