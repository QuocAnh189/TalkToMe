import { useState } from 'react'
import Input from '@ui/Input'
import Button from '@ui/Button'
import { BiUser, BiEnvelope, BiPhone, BiPencil } from 'react-icons/bi'
import Avatar from '@ui/Avatar'

const ProfileInfo = () => {
  const [formData, setFormData] = useState({
    name: 'John Doe',
    email: 'john.doe@example.com',
    phone: '+1234567890',
    bio: 'Software Developer passionate about creating beautiful and functional applications.',
  })

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // Handle form submission
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-8">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="space-y-6">
          <div>
            <Input
              label="Full Name"
              value={formData.name}
              onChange={(e) => setFormData({ ...formData, name: e.target.value })}
              startIcon={<BiUser className="w-5 h-5 text-gray-400" />}
              className="bg-white"
            />
          </div>
          <div>
            <Input
              label="Email"
              type="email"
              value={formData.email}
              onChange={(e) => setFormData({ ...formData, email: e.target.value })}
              startIcon={<BiEnvelope className="w-5 h-5 text-gray-400" />}
              className="bg-white"
            />
          </div>
          <div>
            <Input
              label="Phone"
              value={formData.phone}
              onChange={(e) => setFormData({ ...formData, phone: e.target.value })}
              startIcon={<BiPhone className="w-5 h-5 text-gray-400" />}
              className="bg-white"
            />
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">Profile Picture</label>
          <div className="relative flex flex-col items-center p-6 bg-white rounded-lg border border-gray-200">
            <Avatar
              src="https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser"
              size="xl"
              className="mb-4 w-32 h-32"
            />
            <div className="flex flex-col items-center gap-2">
              <input
                type="file"
                id="avatar"
                className="hidden"
                accept="image/*"
                onChange={(e) => {
                  const file = e.target.files?.[0]
                  if (file) {
                    // Handle avatar upload
                    console.log('Avatar file:', file)
                  }
                }}
              />
              <button
                type="button"
                onClick={() => document.getElementById('avatar')?.click()}
                className="btn btn-primary btn-sm gap-2"
              >
                <BiPencil className="w-4 h-4" />
                Change Avatar
              </button>
              <p className="text-sm text-gray-500">Recommended: Square image, max 2MB</p>
            </div>
          </div>
        </div>
      </div>

      <div className="flex justify-end">
        <Button type="submit" variant="primary" className="px-6 py-2.5 shadow-md hover:shadow-lg transition-all">
          Save Changes
        </Button>
      </div>
    </form>
  )
}

export default ProfileInfo
