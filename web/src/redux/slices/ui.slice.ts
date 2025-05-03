import { createSlice, PayloadAction } from '@reduxjs/toolkit'

interface UIState {
  sidebarOpen: boolean
  theme: 'light' | 'dark'
  notifications: {
    id: string
    message: string
    type: 'success' | 'error' | 'info' | 'warning'
  }[]
}

const initialState: UIState = {
  sidebarOpen: true,
  theme: 'light',
  notifications: []
}

const uiSlice = createSlice({
  name: 'ui',
  initialState,
  reducers: {
    toggleSidebar: (state) => {
      state.sidebarOpen = !state.sidebarOpen
    },
    setTheme: (state, action: PayloadAction<'light' | 'dark'>) => {
      state.theme = action.payload
    },
    addNotification: (state, action: PayloadAction<Omit<UIState['notifications'][0], 'id'>>) => {
      state.notifications.push({
        id: Date.now().toString(),
        ...action.payload
      })
    },
    removeNotification: (state, action: PayloadAction<string>) => {
      state.notifications = state.notifications.filter(n => n.id !== action.payload)
    }
  }
})

export const {
  toggleSidebar,
  setTheme,
  addNotification,
  removeNotification
} = uiSlice.actions
export default uiSlice.reducer