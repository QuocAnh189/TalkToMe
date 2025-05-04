export const mockConversations = [
  // Direct Messages
  {
    id: 1,
    name: "John Doe",
    description: "Personal chat with John",
    is_group: false,
    is_user: true,
    owner_id: 1,
    members: [
      {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      {
        id: "2",
        name: "John Doe",
        email: "john@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=John"
      }
    ],
    user_ids: [1, 2],
    created_at: "2024-01-15T10:00:00Z",
    updated_at: "2024-01-15T10:00:00Z",
    last_message: {
      id: 1,
      content: "Hey, how are you?",
      sender: {
        id: "2",
        name: "John Doe",
        email: "john@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=John"
      },
      created_at: "2024-01-15T10:00:00Z"
    }
  },
  {
    id: 2,
    name: "Emma Wilson",
    description: "Chat with Emma",
    is_group: false,
    is_user: true,
    owner_id: 1,
    members: [
      {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      {
        id: "3",
        name: "Emma Wilson",
        email: "emma@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Emma"
      }
    ],
    user_ids: [1, 3],
    created_at: "2024-01-15T09:30:00Z",
    updated_at: "2024-01-15T09:30:00Z",
    last_message: {
      id: 2,
      content: "Can we discuss the project tomorrow?",
      sender: {
        id: "3",
        name: "Emma Wilson",
        email: "emma@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Emma"
      },
      created_at: "2024-01-15T09:30:00Z"
    }
  },
  {
    id: 3,
    name: "Michael Chen",
    description: "Chat with Michael",
    is_group: false,
    is_user: true,
    owner_id: 1,
    members: [
      {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      {
        id: "4",
        name: "Michael Chen",
        email: "michael@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Michael"
      }
    ],
    user_ids: [1, 4],
    created_at: "2024-01-14T15:20:00Z",
    updated_at: "2024-01-14T15:20:00Z",
    last_message: {
      id: 3,
      content: "Thanks for your help!",
      sender: {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      created_at: "2024-01-14T15:20:00Z"
    }
  },
  {
    id: 4,
    name: "Sarah Johnson",
    description: "Chat with Sarah",
    is_group: false,
    is_user: true,
    owner_id: 1,
    members: [
      {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      {
        id: "5",
        name: "Sarah Johnson",
        email: "sarah@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah"
      }
    ],
    user_ids: [1, 5],
    created_at: "2024-01-14T11:45:00Z",
    updated_at: "2024-01-14T11:45:00Z",
    last_message: {
      id: 4,
      content: "See you at the meeting!",
      sender: {
        id: "5",
        name: "Sarah Johnson",
        email: "sarah@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah"
      },
      created_at: "2024-01-14T11:45:00Z"
    }
  },
  {
    id: 5,
    name: "David Kim",
    description: "Chat with David",
    is_group: false,
    is_user: true,
    owner_id: 1,
    members: [
      {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      {
        id: "6",
        name: "David Kim",
        email: "david@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=David"
      }
    ],
    user_ids: [1, 6],
    created_at: "2024-01-13T16:15:00Z",
    updated_at: "2024-01-13T16:15:00Z",
    last_message: {
      id: 5,
      content: "The documents are ready for review",
      sender: {
        id: "6",
        name: "David Kim",
        email: "david@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=David"
      },
      created_at: "2024-01-13T16:15:00Z"
    }
  },
  // Group Chats
  {
    id: 6,
    name: "Project Team",
    description: "Main project discussion group",
    is_group: true,
    is_user: false,
    owner_id: 1,
    members: [
      {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      {
        id: "2",
        name: "John Doe",
        email: "john@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=John"
      },
      {
        id: "3",
        name: "Emma Wilson",
        email: "emma@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Emma"
      }
    ],
    user_ids: [1, 2, 3],
    created_at: "2024-01-15T08:00:00Z",
    updated_at: "2024-01-15T08:00:00Z",
    last_message: {
      id: 6,
      content: "Next sprint planning tomorrow at 10 AM",
      sender: {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      created_at: "2024-01-15T08:00:00Z"
    }
  },
  {
    id: 7,
    name: "Design Team",
    description: "Design discussions and updates",
    is_group: true,
    is_user: false,
    owner_id: 3,
    members: [
      {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      {
        id: "4",
        name: "Michael Chen",
        email: "michael@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Michael"
      },
      {
        id: "5",
        name: "Sarah Johnson",
        email: "sarah@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah"
      }
    ],
    user_ids: [1, 4, 5],
    created_at: "2024-01-14T13:30:00Z",
    updated_at: "2024-01-14T13:30:00Z",
    last_message: {
      id: 7,
      content: "New design mockups uploaded",
      sender: {
        id: "4",
        name: "Michael Chen",
        email: "michael@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Michael"
      },
      created_at: "2024-01-14T13:30:00Z"
    }
  },
  {
    id: 8,
    name: "Marketing Team",
    description: "Marketing strategy and campaigns",
    is_group: true,
    is_user: false,
    owner_id: 5,
    members: [
      {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      {
        id: "5",
        name: "Sarah Johnson",
        email: "sarah@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah"
      },
      {
        id: "6",
        name: "David Kim",
        email: "david@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=David"
      }
    ],
    user_ids: [1, 5, 6],
    created_at: "2024-01-13T14:20:00Z",
    updated_at: "2024-01-13T14:20:00Z",
    last_message: {
      id: 8,
      content: "Q1 campaign results are in",
      sender: {
        id: "5",
        name: "Sarah Johnson",
        email: "sarah@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah"
      },
      created_at: "2024-01-13T14:20:00Z"
    }
  },
  {
    id: 9,
    name: "Development Team",
    description: "Technical discussions and updates",
    is_group: true,
    is_user: false,
    owner_id: 2,
    members: [
      {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      {
        id: "2",
        name: "John Doe",
        email: "john@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=John"
      },
      {
        id: "4",
        name: "Michael Chen",
        email: "michael@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Michael"
      }
    ],
    user_ids: [1, 2, 4],
    created_at: "2024-01-12T16:45:00Z",
    updated_at: "2024-01-12T16:45:00Z",
    last_message: {
      id: 9,
      content: "Code review needed for the new feature",
      sender: {
        id: "2",
        name: "John Doe",
        email: "john@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=John"
      },
      created_at: "2024-01-12T16:45:00Z"
    }
  },
  {
    id: 10,
    name: "General",
    description: "Company-wide announcements",
    is_group: true,
    is_user: false,
    owner_id: 1,
    members: [
      {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      {
        id: "2",
        name: "John Doe",
        email: "john@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=John"
      },
      {
        id: "3",
        name: "Emma Wilson",
        email: "emma@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Emma"
      },
      {
        id: "4",
        name: "Michael Chen",
        email: "michael@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Michael"
      },
      {
        id: "5",
        name: "Sarah Johnson",
        email: "sarah@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah"
      },
      {
        id: "6",
        name: "David Kim",
        email: "david@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=David"
      }
    ],
    user_ids: [1, 2, 3, 4, 5, 6],
    created_at: "2024-01-11T09:00:00Z",
    updated_at: "2024-01-11T09:00:00Z",
    last_message: {
      id: 10,
      content: "Monthly all-hands meeting this Friday",
      sender: {
        id: "1",
        name: "Current User",
        email: "current@example.com",
        avatar_url: "https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
      },
      created_at: "2024-01-11T09:00:00Z"
    }
  }
];

export default mockConversations;