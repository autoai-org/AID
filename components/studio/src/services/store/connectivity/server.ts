import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { RootState } from '../store'

interface ConnectivityState {
    isLoading: boolean,
    server: string
}

const initialState: ConnectivityState = {
    isLoading: false,
    server: "",
}

export const connectivitySlice = createSlice({
    name:"connectivity",
    initialState,
    reducers: {
        setServer: (state, action: PayloadAction<string>) => {
            state.server = action.payload
        },
        setIsLoading: (state, action: PayloadAction<boolean>) => {
            state.isLoading = action.payload
        }
    }
})

export const { setServer } = connectivitySlice.actions

export const getServer = (state:RootState) => state.connectivity.server

export default connectivitySlice.reducer