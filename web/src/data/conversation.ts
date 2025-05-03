const mockConversations = [
    {
        id: '1',
        partner: {
            id: 'u1',
            name: 'John Doe',
            avatarURL: 'https://i.pravatar.cc/150?img=1',
            isOnline: true
        },
        lastMessage: {
            id: 'm1',
            message: 'Hey, how are you doing?',
            createdAt: '2024-01-20T10:30:00Z'
        },
        unreadCount: 2
    },
    {
        id: '2',
        partner: {
            id: 'u2',
            name: 'Alice Smith',
            avatarURL: 'https://i.pravatar.cc/150?img=2',
            isOnline: false
        },
        lastMessage: {
            id: 'm2',
            message: 'The project deadline is tomorrow!',
            createdAt: '2024-01-20T09:15:00Z'
        },
        unreadCount: 0
    },
    {
        id: '3',
        partner: {
            id: 'u3',
            name: 'Bob Johnson',
            avatarURL: 'https://i.pravatar.cc/150?img=3',
            isOnline: true
        },
        lastMessage: {
            id: 'm3',
            message: 'Let\'s meet for coffee',
            createdAt: '2024-01-19T18:45:00Z'
        },
        unreadCount: 1
    },
    {
        id: '4',
        partner: {
        id: 'u4',
            name: 'Bob Johnson',
            avatarURL: 'https://i.pravatar.cc/150?img=3',
            isOnline: true
        },
        lastMessage: {
            id: 'm4',
            message: 'Let\'s meet for coffee',
            createdAt: '2024-01-19T18:45:00Z'
        },
        unreadCount: 1
    },
    {
        id: '5',
        partner: {
            id: 'u3',
            name: 'Bob Johnson',
            avatarURL: 'https://i.pravatar.cc/150?img=5',
            isOnline: true
        },
        lastMessage: {
            id: 'm5',
            message: 'Let\'s meet for coffee',
            createdAt: '2024-01-19T18:45:00Z'
        },
        unreadCount: 1
    },
    {
        id: '6',
        partner: {
            id: 'u6',
            name: 'Bob Johnson',
            avatarURL: 'https://i.pravatar.cc/150?img=6',
            isOnline: true
        },
        lastMessage: {
            id: 'm6',
            message: 'Let\'s meet for coffee',
            createdAt: '2024-01-19T18:45:00Z'
        },
        unreadCount: 1
    },
    {
        id: '7',
        partner: {
            id: 'u7',
            name: 'Bob Johnson',
            avatarURL: 'https://i.pravatar.cc/150?img=7',
            isOnline: true
        },
        lastMessage: {
            id: 'm3',
            message: 'Let\'s meet for coffee',
            createdAt: '2024-01-19T18:45:00Z'
        },
        unreadCount: 1
    },
    {
        id: '7',
        partner: {
            id: 'u7',
            name: 'Bob Johnson',
            avatarURL: 'https://i.pravatar.cc/150?img=8',
            isOnline: true
        },
        lastMessage: {
            id: 'm3',
            message: 'Let\'s meet for coffee',
            createdAt: '2024-01-19T18:45:00Z'
        },
        unreadCount: 1
    }
]

// Mock data for groups
const mockGroups = [
    {
        id: 'g1',
        name: 'Project Team',
        description: 'Team discussion group',
        avatarURL: 'https://i.pravatar.cc/150?img=4',
        lastMessage: {
            id: 'm4',
            message: 'Meeting at 3 PM today',
            createdAt: '2024-01-20T08:00:00Z'
        },
        unreadCount: 5
    },
    {
        id: 'g2',
        name: 'Family Group',
        description: 'Family chat',
        avatarURL: 'https://i.pravatar.cc/150?img=5',
        lastMessage: {
            id: 'm5',
            message: 'Who\'s coming to dinner?',
            createdAt: '2024-01-19T20:30:00Z'
        },
        unreadCount: 0
    }
]

export { mockConversations, mockGroups };