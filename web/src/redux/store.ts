import { configureStore } from '@reduxjs/toolkit'

//slices
import authReducer from './slices/auth.slice'
import chatReducer from './slices/chat.slice'
import uiReducer from './slices/ui.slice'

//services
import { apiAuth } from './services/auth'
import { apiUser } from './services/user'
import { apiMessage } from './services/message'
import { apiFriend } from './services/friend'
import { apiConversation } from './services/conversation'
import { apiGroup } from './services/group'
import { apiNotification } from './services/notification'

const store = configureStore({
  reducer: {
    auth: authReducer,
    chat: chatReducer,
    ui: uiReducer,
    
    [apiAuth.reducerPath]: apiAuth.reducer,
    [apiUser.reducerPath]: apiUser.reducer,
    [apiMessage.reducerPath]: apiMessage.reducer,
    [apiFriend.reducerPath]: apiFriend.reducer,
    [apiConversation.reducerPath]: apiConversation.reducer,
    [apiGroup.reducerPath]: apiGroup.reducer,
    [apiNotification.reducerPath]: apiNotification.reducer,
  },

  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
      immutableCheck: false,
    }).concat([
      apiAuth.middleware,
      apiUser.middleware,
      apiMessage.middleware,
      apiFriend.middleware,
      apiConversation.middleware,
      apiGroup.middleware,
      apiNotification.middleware,
    ]),
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch

export default store
