import { configureStore } from "@reduxjs/toolkit";
import userReducer from "./userSlice";
import threadReducer from "./threadSlice";
import tagReducer from "./tagSlice";
import commentReducer from "./commentsSlice";

const store = configureStore({
  reducer: {
    user: userReducer,
    thread: threadReducer,
    tag: tagReducer,
    comment: commentReducer
  },
})

export type RootState = ReturnType<typeof store.getState>; //for typescript
export type AppDispatch = typeof store.dispatch;

export default store;