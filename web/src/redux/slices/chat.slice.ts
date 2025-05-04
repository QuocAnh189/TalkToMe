import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { IMessage } from '@interfaces/message'
import { IConversation } from '@interfaces/conversation'
import { IGroup } from '@interfaces/group'

export const ChatSliceKey = 'chat'

interface ChatState {
  activeChat: {
    id: string
    type: 'conversation' | 'group'
  } | null
  messages: IMessage[]
  conversations: IConversation[]
  groups: IGroup[]
  loading: boolean
  error: string | null
}

const initialState: ChatState = {
  activeChat: null,
  messages: [],
  conversations: [],
  groups: [],
  loading: false,
  error: null
}

const chatSlice = createSlice({
  name: ChatSliceKey,
  initialState,
  reducers: {
    setActiveChat: (state, action: PayloadAction<ChatState['activeChat']>) => {
      state.activeChat = action.payload
    },
    setMessages: (state, action: PayloadAction<IMessage[]>) => {
      state.messages = action.payload
    },
    addMessage: (state, action: PayloadAction<IMessage>) => {
      state.messages.push(action.payload)
    },
    setConversations: (state, action: PayloadAction<IConversation[]>) => {
      state.conversations = action.payload
    },
    setGroups: (state, action: PayloadAction<IGroup[]>) => {
      state.groups = action.payload
    },
    setChatLoading: (state, action: PayloadAction<boolean>) => {
      state.loading = action.payload
    },
    setChatError: (state, action: PayloadAction<string | null>) => {
      state.error = action.payload
    }
  }
})

export const {
  setActiveChat,
  setMessages,
  addMessage,
  setConversations,
  setGroups,
  setChatLoading,
  setChatError
} = chatSlice.actions
export default chatSlice.reducer