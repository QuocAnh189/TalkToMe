import { createApi } from '@reduxjs/toolkit/query/react'
import baseQueryWithReauth from '@redux/interceptor/baseQueryWithReauth'

//interfaces
import { IUser } from '@interfaces/user'

export const apiFriend = createApi({
  reducerPath: 'apiFriend',
  baseQuery: baseQueryWithReauth,
  tagTypes: ['Friend'],
  endpoints: (builder) => ({
    listFriends: builder.query<IUser[], { page?: number; limit?: number }>({
      query: (params) => ({
        url: '/friends',
        method: 'GET',
        params,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ['Friend'],
    }),

    addFriend: builder.mutation<void, string>({
      query: (friendId) => ({
        url: '/friends/add',
        method: 'POST',
        body: { friendId },
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Friend'],
    }),

    removeFriend: builder.mutation<void, string>({
      query: (friendId) => ({
        url: '/friends/remove',
        method: 'DELETE',
        body: { friendId },
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Friend'],
    }),
  }),
})

export const {
  useListFriendsQuery,
  useAddFriendMutation,
  useRemoveFriendMutation,
} = apiFriend