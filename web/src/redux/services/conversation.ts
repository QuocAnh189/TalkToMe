import { createApi } from '@reduxjs/toolkit/query/react'
import baseQueryWithReauth from '@redux/interceptor/baseQueryWithReauth'

//interfaces
import { IConversation } from '@interfaces/conversation'

export const apiConversation = createApi({
  reducerPath: 'apiConversation',
  baseQuery: baseQueryWithReauth,
  tagTypes: ['Conversation'],
  endpoints: (builder) => ({
    listConversations: builder.query<IConversation[], { page?: number; limit?: number }>({
      query: (params) => ({
        url: '/conversations',
        method: 'GET',
        params,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ['Conversation'],
    }),

    getConversation: builder.query<IConversation, string>({
      query: (id) => ({
        url: `/conversations/${id}`,
        method: 'GET',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),

    createConversation: builder.mutation<IConversation, any>({
      query: (data) => ({
        url: '/conversations',
        method: 'POST',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Conversation'],
    }),

    deleteConversation: builder.mutation<void, string>({
      query: (id) => ({
        url: `/conversations/${id}`,
        method: 'DELETE',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Conversation'],
    }),
  }),
})

export const {
  useListConversationsQuery,
  useGetConversationQuery,
  useCreateConversationMutation,
  useDeleteConversationMutation,
} = apiConversation