import { useState, useRef } from 'react'
import Modal from '@ui/Modal'
import Avatar from '@ui/Avatar'

interface ChangeGroupAvatarModalProps {
  isOpen: boolean
  onClose: () => void
  onConfirm: (file: File) => void
  currentAvatar: string
}

const ChangeGroupAvatarModal = ({ isOpen, onClose, onConfirm, currentAvatar }: ChangeGroupAvatarModalProps) => {
  const [previewUrl, setPreviewUrl] = useState<string>(currentAvatar)
  const [selectedFile, setSelectedFile] = useState<File | null>(null)
  const fileInputRef = useRef<HTMLInputElement>(null)

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0]
    if (file) {
      setSelectedFile(file)
      const reader = new FileReader()
      reader.onloadend = () => {
        setPreviewUrl(reader.result as string)
      }
      reader.readAsDataURL(file)
    }
  }

  const handleSubmit = () => {
    if (selectedFile) {
      onConfirm(selectedFile)
      setSelectedFile(null)
      setPreviewUrl(currentAvatar)
    }
  }

  const handleSelectFile = () => {
    fileInputRef.current?.click()
  }

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Change Group Avatar">
      <div className="py-4">
        <div className="flex flex-col items-center gap-4">
          <Avatar src={previewUrl} size="lg" />

          <input type="file" ref={fileInputRef} className="hidden" accept="image/*" onChange={handleFileChange} />

          <button className="btn btn-primary" onClick={handleSelectFile}>
            Choose Image
          </button>

          {selectedFile && <p className="text-sm text-base-content/60">Selected: {selectedFile.name}</p>}
        </div>

        <div className="flex justify-end gap-2 mt-6">
          <button
            className="btn btn-ghost"
            onClick={() => {
              onClose()
              setPreviewUrl(currentAvatar)
              setSelectedFile(null)
            }}
          >
            Cancel
          </button>
          <button className="btn btn-primary" onClick={handleSubmit} disabled={!selectedFile}>
            Save
          </button>
        </div>
      </div>
    </Modal>
  )
}

export default ChangeGroupAvatarModal
