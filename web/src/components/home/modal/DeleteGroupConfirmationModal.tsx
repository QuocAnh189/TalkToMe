import Modal from '@ui/Modal'

interface DeleteConfirmationModalProps {
  isOpen: boolean
  onClose: () => void
  onConfirm: () => void
}

const DeleteGroupConfirmationModal = ({ isOpen, onClose, onConfirm }: DeleteConfirmationModalProps) => {
  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Delete Group">
      <div className="py-4">
        <p className="text-base-content/80">Are you sure you want to delete this group ?</p>
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

export default DeleteGroupConfirmationModal
