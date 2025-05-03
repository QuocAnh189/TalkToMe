import Modal from '@ui/Modal'

interface ExitGroupConfirmationModalProps {
  isOpen: boolean
  onClose: () => void
  onConfirm: () => void
}

const ExitGroupConfirmationModal = ({ isOpen, onClose, onConfirm }: ExitGroupConfirmationModalProps) => {
  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Exit Group">
      <div className="py-4">
        <p className="text-base-content/80">Are you sure you want to exit this group?</p>
        <div className="flex justify-end gap-2 mt-6">
          <button className="btn btn-ghost" onClick={onClose}>
            Cancel
          </button>
          <button className="btn btn-error" onClick={onConfirm}>
            Exit
          </button>
        </div>
      </div>
    </Modal>
  )
}

export default ExitGroupConfirmationModal
