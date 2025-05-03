import { createApi } from '@reduxjs/toolkit/query/react'
import baseQueryWithReauth from '@redux/interceptor/baseQueryWithReauth'

//interfaces
import { INotification } from '@interfaces/notification'

export const apiNotification = createApi({
  reducerPath: 'apiNotification',
  baseQuery: baseQueryWithReauth,
  tagTypes: ['Notification'],
  endpoints: (builder) => ({
    listNotifications: builder.query<INotification[], { page?: number; limit?: number }>({
      query: (params) => ({
        url: '/notifications',
        method: 'GET',
        params,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ['Notification'],
    }),

    countUnread: builder.query<number, void>({
      query: () => ({
        url: '/notifications/unread/count',
        method: 'GET',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ['Notification'],
    }),

    markAsRead: builder.mutation<void, string>({
      query: (id) => ({
        url: `/notifications/read/${id}`,
        method: 'PUT',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Notification'],
    }),

    markAllAsRead: builder.mutation<void, void>({
      query: () => ({
        url: '/notifications/read-all',
        method: 'PUT',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Notification'],
    }),

    deleteNotification: builder.mutation<void, string>({
      query: (id) => ({
        url: `/notifications/${id}`,
        method: 'DELETE',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Notification'],
    }),
  }),
})

export const {
  useListNotificationsQuery,
  useCountUnreadQuery,
  useMarkAsReadMutation,
  useMarkAllAsReadMutation,
  useDeleteNotificationMutation,
} = apiNotification