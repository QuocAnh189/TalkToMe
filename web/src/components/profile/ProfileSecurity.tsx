import { useState } from 'react'
import Input from '@ui/Input'
import Button from '@ui/Button'
import { BiLock, BiShield } from 'react-icons/bi'

const ProfileSecurity = () => {
  const [formData, setFormData] = useState({
    currentPassword: '',
    newPassword: '',
    confirmPassword: '',
  })

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // Handle password change
  }

  return (
    <form onSubmit={handleSubmit} className="max-w-md mx-auto space-y-8">
      <div className="bg-blue-50 border border-blue-100 rounded-lg p-4 flex items-start gap-3">
        <BiShield className="w-6 h-6 text-blue-500 mt-0.5" />
        <div>
          <h3 className="font-medium text-blue-900">Password Security</h3>
          <p className="text-sm text-blue-700 mt-1">Choose a strong password and don't reuse it for other accounts.</p>
        </div>
      </div>

      <div className="space-y-6">
        <Input
          label="Current Password"
          type="password"
          value={formData.currentPassword}
          onChange={(e) => setFormData({ ...formData, currentPassword: e.target.value })}
          startIcon={<BiLock className="w-5 h-5 text-gray-400" />}
          className="bg-white"
        />
        <Input
          label="New Password"
          type="password"
          value={formData.newPassword}
          onChange={(e) => setFormData({ ...formData, newPassword: e.target.value })}
          startIcon={<BiLock className="w-5 h-5 text-gray-400" />}
          className="bg-white"
        />
        <Input
          label="Confirm New Password"
          type="password"
          value={formData.confirmPassword}
          onChange={(e) => setFormData({ ...formData, confirmPassword: e.target.value })}
          startIcon={<BiLock className="w-5 h-5 text-gray-400" />}
          className="bg-white"
        />
      </div>

      <div className="flex justify-end">
        <Button type="submit" variant="primary" className="px-6 py-2.5 shadow-md hover:shadow-lg transition-all">
          Update Password
        </Button>
      </div>
    </form>
  )
}

export default ProfileSecurity
