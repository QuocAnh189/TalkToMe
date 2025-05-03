import { createApi } from '@reduxjs/toolkit/query/react'
import baseQueryWithReauth from '@redux/interceptor/baseQueryWithReauth'

//interfaces
import { IMessage, SendMessageRequest, UpdateMessageRequest } from '@interfaces/message'

export const apiMessage = createApi({
  reducerPath: 'apiMessage',
  baseQuery: baseQueryWithReauth,
  tagTypes: ['Message'],
  endpoints: (builder) => ({
    sendMessage: builder.mutation<IMessage, SendMessageRequest>({
      query: (data) => ({
        url: '/messages',
        method: 'POST',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Message'],
    }),

    getMessage: builder.query<IMessage, string>({
      query: (messageId) => ({
        url: `/messages/${messageId}`,
        method: 'GET',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),

    getGroupMessages: builder.query<IMessage[], { groupId: string; page?: number; limit?: number }>({
      query: ({ groupId, ...params }) => ({
        url: `/messages/group/${groupId}`,
        method: 'GET',
        params,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ['Message'],
    }),

    getConversationMessages: builder.query<IMessage[], { conversationId: string; page?: number; limit?: number }>({
      query: ({ conversationId, ...params }) => ({
        url: `/messages/conversation/${conversationId}`,
        method: 'GET',
        params,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ['Message'],
    }),

    updateMessage: builder.mutation<IMessage, { messageId: string; data: UpdateMessageRequest }>({
      query: ({ messageId, data }) => ({
        url: `/messages/${messageId}`,
        method: 'PUT',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Message'],
    }),

    deleteMessage: builder.mutation<void, string>({
      query: (messageId) => ({
        url: `/messages/${messageId}`,
        method: 'DELETE',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Message'],
    }),
  }),
})

export const {
  useSendMessageMutation,
  useGetMessageQuery,
  useGetGroupMessagesQuery,
  useGetConversationMessagesQuery,
  useUpdateMessageMutation,
  useDeleteMessageMutation,
} = apiMessage