import Modal from '@ui/Modal'

interface DeleteChatConfirmationModalProps {
  isOpen: boolean
  onClose: () => void
  onConfirm: () => void
  friendName: string
}

const DeleteChatConfirmationModal = ({ isOpen, onClose, onConfirm, friendName }: DeleteChatConfirmationModalProps) => {
  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Delete Chat">
      <div className="py-4">
        <p className="text-base-content/80">
          Are you sure you want to delete your chat with <span className="font-medium">{friendName}</span> ?
        </p>

        <div className="flex justify-end gap-2 mt-6">
          <button className="btn btn-ghost" onClick={onClose}>
            Cancel
          </button>
          <button className="btn btn-error" onClick={onConfirm}>
            Delete
          </button>
        </div>
      </div>
    </Modal>
  )
}

export default DeleteChatConfirmationModal
