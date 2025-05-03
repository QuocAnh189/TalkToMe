import { useState } from 'react'
import Modal from '@ui/Modal'
import Input from '@ui/Input'

interface ChangeGroupNameModalProps {
  isOpen: boolean
  onClose: () => void
  onConfirm: (name: string) => void
  currentName: string
}

const ChangeGroupNameModal = ({ isOpen, onClose, onConfirm, currentName }: ChangeGroupNameModalProps) => {
  const [name, setName] = useState(currentName)

  const handleSubmit = () => {
    if (name.trim() && name.trim() !== currentName) {
      onConfirm(name.trim())
      setName(currentName)
    }
  }

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Change Group Name">
      <div className="py-4">
        <Input
          label="Group Name"
          placeholder="Enter group name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <div className="flex justify-end gap-2 mt-6">
          <button className="btn btn-ghost" onClick={onClose}>
            Cancel
          </button>
          <button
            className="btn btn-primary"
            onClick={handleSubmit}
            disabled={!name.trim() || name.trim() === currentName}
          >
            Save
          </button>
        </div>
      </div>
    </Modal>
  )
}

export default ChangeGroupNameModal
