import { cn } from '@utils/cn'

interface ModalProps {
  isOpen: boolean
  onClose: () => void
  children: React.ReactNode
  title?: string
  className?: string
  showCloseButton?: boolean
}

const Modal = ({ isOpen, onClose, children, title, className, showCloseButton = true }: ModalProps) => {
  if (!isOpen) return null

  return (
    <div className="modal modal-open">
      <div className={cn('modal-box relative', className)}>
        {showCloseButton && (
          <button className="btn btn-sm btn-circle absolute right-2 top-2" onClick={onClose}>
            ✕
          </button>
        )}
        {title && <h3 className="font-bold text-lg mb-4">{title}</h3>}
        {children}
      </div>
      <div className="modal-backdrop" onClick={onClose}></div>
    </div>
  )
}

export default Modal
