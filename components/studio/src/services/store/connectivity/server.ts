import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { RootState } from '../store'
import {connect, ConnectedProps} from 'react-redux';
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

const mapState = (state: RootState) => ({
    isLoading: state.connectivity.isLoading
})

const mapDispatch = {
    toggleLoading: () => ({type: "TOGGLE_IS_LOADING"})
}
export const connector = connect(mapState, mapDispatch)

type PropsFromRedux = ConnectedProps<typeof connector>

export const { setServer, setIsLoading } = connectivitySlice.actions

export const getServer = (state:RootState) => state.connectivity.server

export type { PropsFromRedux }

export default connectivitySlice.reducer