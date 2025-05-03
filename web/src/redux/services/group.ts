import { createApi } from '@reduxjs/toolkit/query/react'
import baseQueryWithReauth from '@redux/interceptor/baseQueryWithReauth'

//interfaces
import { IGroup, CreateGroupRequest, UpdateGroupRequest } from '@interfaces/group'
import { IUser } from '@interfaces/user'

export const apiGroup = createApi({
  reducerPath: 'apiGroup',
  baseQuery: baseQueryWithReauth,
  tagTypes: ['Group'],
  endpoints: (builder) => ({
    listGroups: builder.query<IGroup[], { page?: number; limit?: number }>({
      query: (params) => ({
        url: '/groups',
        method: 'GET',
        params,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ['Group'],
    }),

    getGroup: builder.query<IGroup, string>({
      query: (id) => ({
        url: `/groups/${id}`,
        method: 'GET',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),

    createGroup: builder.mutation<IGroup, CreateGroupRequest>({
      query: (data) => ({
        url: '/groups',
        method: 'POST',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Group'],
    }),

    updateGroup: builder.mutation<IGroup, { groupId: string; data: UpdateGroupRequest }>({
      query: ({ groupId, data }) => ({
        url: `/groups/${groupId}`,
        method: 'PUT',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Group'],
    }),

    deleteGroup: builder.mutation<void, string>({
      query: (id) => ({
        url: `/groups/${id}`,
        method: 'DELETE',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Group'],
    }),

    addMember: builder.mutation<void, { groupId: string; userId: string }>({
      query: ({ groupId, userId }) => ({
        url: `/groups/${groupId}/members`,
        method: 'POST',
        body: { userId },
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Group'],
    }),

    removeMember: builder.mutation<void, { groupId: string; userId: string }>({
      query: ({ groupId, userId }) => ({
        url: `/groups/${groupId}/members/${userId}`,
        method: 'DELETE',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Group'],
    }),

    listMembers: builder.query<IUser[], string>({
      query: (groupId) => ({
        url: `/groups/${groupId}/members`,
        method: 'GET',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),
  }),
})

export const {
  useListGroupsQuery,
  useGetGroupQuery,
  useCreateGroupMutation,
  useUpdateGroupMutation,
  useDeleteGroupMutation,
  useAddMemberMutation,
  useRemoveMemberMutation,
  useListMembersQuery,
} = apiGroup