import { useState } from 'react'

import ProfileHeader from '@components/profile/ProfileHeader'
import ProfileInfo from '@components/profile/ProfileInfo'
import ProfileSecurity from '@components/profile/ProfileSecurity'
import ProtectedLayout from '@components/layout/protected'
import { BiUser, BiLock } from 'react-icons/bi'

const Profile = () => {
  const [activeTab, setActiveTab] = useState<'info' | 'security'>('info')

  return (
    <ProtectedLayout>
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-50 p-8">
        <div className="max-w-4xl mx-auto">
          <div className="bg-white rounded-2xl shadow-xl overflow-hidden">
            <ProfileHeader />

            <div className="p-8">
              <div className="flex gap-6 mb-8">
                <button
                  className={`flex items-center gap-2 py-2 px-4 rounded-lg transition-all ${
                    activeTab === 'info' ? 'bg-primary text-white shadow-md' : 'hover:bg-gray-100'
                  }`}
                  onClick={() => setActiveTab('info')}
                >
                  <BiUser className="w-5 h-5" />
                  <span>Personal Info</span>
                </button>
                <button
                  className={`flex items-center gap-2 py-2 px-4 rounded-lg transition-all ${
                    activeTab === 'security' ? 'bg-primary text-white shadow-md' : 'hover:bg-gray-100'
                  }`}
                  onClick={() => setActiveTab('security')}
                >
                  <BiLock className="w-5 h-5" />
                  <span>Security</span>
                </button>
              </div>

              <div className="bg-gray-50 rounded-xl p-6">
                {activeTab === 'info' ? <ProfileInfo /> : <ProfileSecurity />}
              </div>
            </div>
          </div>
        </div>
      </div>
    </ProtectedLayout>
  )
}

export default Profile
