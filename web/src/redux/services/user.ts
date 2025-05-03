import { createApi } from '@reduxjs/toolkit/query/react'
import baseQueryWithReauth from '@redux/interceptor/baseQueryWithReauth'

//interfaces
import { IUser } from '@interfaces/user'

export const apiUser = createApi({
  reducerPath: 'apiUser',
  baseQuery: baseQueryWithReauth,
  tagTypes: ['User'],
  endpoints: (builder) => ({
    getProfile: builder.query<IUser, void>({
      query: () => ({
        url: '/users/me',
        method: 'GET',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),

    updateProfile: builder.mutation<IUser, FormData>({
      query: (data) => ({
        url: '/users/me',
        method: 'PUT',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['User'],
    }),

    getUserById: builder.query<IUser, string>({
      query: (id) => ({
        url: `/users/${id}`,
        method: 'GET',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),

    searchUsers: builder.query<IUser[], { search?: string; page?: number; limit?: number }>({
      query: (params) => ({
        url: '/users',
        method: 'GET',
        params,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),
  }),
})

export const { 
  useGetProfileQuery, 
  useUpdateProfileMutation, 
  useGetUserByIdQuery,
  useSearchUsersQuery 
} = apiUser