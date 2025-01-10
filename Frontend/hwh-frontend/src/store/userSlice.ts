import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface User {
  userId: string;
  username: string;
  role: string;
  createdAt: string;
}

interface UserState {
  user: User | null;
}

const initialState: UserState = {
  user: null
}

export const fetchUserData = createAsyncThunk("user/fetchUserData", async (token: string) => {
  try {
    const response = await fetch(`http://localhost:8080/account/user`, {
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${token}`
      }
    });
    const content = await response.json();
    return content.data.user;
  } catch (err) {
    console.log("Error:", err);
  }
});

export const userSlice = createSlice({
  name: "user",
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
    clearUser: (state) => {
      state.user = null;
    }
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchUserData.pending, (state) => {
        // state.isLoading = true;
        console.log("Getting user data pending");
      })
      .addCase(fetchUserData.fulfilled, (state, action) => {
        // state.isLoading = false;
        console.log("Getting user data fulfilled");
        state.user = action.payload;
      })
      .addCase(fetchUserData.rejected, (state, action) => {
        // state.isLoading = false;
        // state.error = action.error.message || "Failed to fetch habits";
        console.log("Getting user data rejected");
      });
  }
});

export const { clearUser } = userSlice.actions;

export default userSlice.reducer;