import { useState, useEffect } from 'react'
import Modal from '@ui/Modal'
import Input from '@ui/Input'
import Avatar from '@ui/Avatar'

interface GroupMember {
  id: string
  name: string
  nickname?: string
  avatarURL: string
}

interface ChangeNicknameModalProps {
  isOpen: boolean
  onClose: () => void
  onConfirm: (memberId: string, nickname: string) => void
  groupMembers: GroupMember[]
  currentUserId: string
}

const ChangeNicknameGroupModal = ({
  isOpen,
  onClose,
  onConfirm,
  groupMembers,
  currentUserId,
}: ChangeNicknameModalProps) => {
  const [selectedMember, setSelectedMember] = useState<string>('')
  const [nickname, setNickname] = useState('')

  useEffect(() => {
    if (isOpen) {
      setSelectedMember(currentUserId)
      const currentMember = groupMembers.find((m) => m.id === currentUserId)
      setNickname(currentMember?.nickname || '')
    }
  }, [isOpen, currentUserId, groupMembers])

  const handleSubmit = () => {
    if (selectedMember && nickname.trim()) {
      onConfirm(selectedMember, nickname.trim())
      setNickname('')
    }
  }

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Change Nickname">
      <div className="py-4">
        <div className="mb-4">
          <label className="font-medium mb-2 block">Select Member</label>
          <div className="space-y-2 max-h-48 overflow-y-auto">
            {groupMembers.map((member) => (
              <div
                key={member.id}
                className={`flex items-center gap-3 p-2 rounded-lg cursor-pointer hover:bg-base-200 transition-colors ${
                  selectedMember === member.id ? 'bg-base-200' : ''
                }`}
                onClick={() => {
                  setSelectedMember(member.id)
                  setNickname(member.nickname || '')
                }}
              >
                <Avatar src={member.avatarURL} size="sm" />
                <div>
                  <div className="font-medium">{member.name}</div>
                  {member.nickname && <div className="text-sm text-base-content/60">Nickname: {member.nickname}</div>}
                </div>
              </div>
            ))}
          </div>
        </div>

        <Input
          label="New Nickname"
          placeholder="Enter new nickname"
          value={nickname}
          onChange={(e) => setNickname(e.target.value)}
        />

        <div className="flex justify-end gap-2 mt-6">
          <button className="btn btn-ghost" onClick={onClose}>
            Cancel
          </button>
          <button className="btn btn-primary" onClick={handleSubmit} disabled={!selectedMember || !nickname.trim()}>
            Save
          </button>
        </div>
      </div>
    </Modal>
  )
}

export default ChangeNicknameGroupModal
