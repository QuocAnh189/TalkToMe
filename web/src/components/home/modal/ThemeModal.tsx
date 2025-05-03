import Modal from '@ui/Modal'

interface Theme {
  id: string
  name: string
  colors: {
    primary: string
    secondary: string
    background: string
    text: string
  }
  preview: string
}

interface ThemeModalProps {
  isOpen: boolean
  onClose: () => void
  onSelectTheme: (themeId: string) => void
  currentThemeId: string
}

const ThemeModal = ({ isOpen, onClose, onSelectTheme, currentThemeId }: ThemeModalProps) => {
  const themes: Theme[] = [
    {
      id: 'default',
      name: 'Default Gray',
      colors: {
        primary: '#e2e8f0',
        secondary: '#e2e8f0',
        background: '#ffffff',
        text: '#000000',
      },
      preview: 'linear-gradient(45deg, #1976d2, #e3f2fd)',
    },
    {
      id: 'blue',
      name: 'Blue Light',
      colors: {
        primary: '#1976d2',
        secondary: '#e3f2fd',
        background: '#ffffff',
        text: '#000000',
      },
      preview: 'linear-gradient(45deg, #1976d2, #e3f2fd)',
    },
    {
      id: 'nature',
      name: 'Nature Green',
      colors: {
        primary: '#2e7d32',
        secondary: '#e8f5e9',
        background: '#ffffff',
        text: '#000000',
      },
      preview: 'linear-gradient(45deg, #2e7d32, #e8f5e9)',
    },
    {
      id: 'sunset',
      name: 'Sunset Orange',
      colors: {
        primary: '#f57c00',
        secondary: '#fff3e0',
        background: '#ffffff',
        text: '#000000',
      },
      preview: 'linear-gradient(45deg, #f57c00, #fff3e0)',
    },
    {
      id: 'purple',
      name: 'Royal Purple',
      colors: {
        primary: '#7b1fa2',
        secondary: '#f3e5f5',
        background: '#ffffff',
        text: '#000000',
      },
      preview: 'linear-gradient(45deg, #7b1fa2, #f3e5f5)',
    },
    {
      id: 'ocean',
      name: 'Ocean Breeze',
      colors: {
        primary: '#0097a7',
        secondary: '#e0f7fa',
        background: '#ffffff',
        text: '#000000',
      },
      preview: 'linear-gradient(45deg, #0097a7, #e0f7fa)',
    },
    {
      id: 'rose',
      name: 'Rose Gold',
      colors: {
        primary: '#c2185b',
        secondary: '#fce4ec',
        background: '#ffffff',
        text: '#000000',
      },
      preview: 'linear-gradient(45deg, #c2185b, #fce4ec)',
    },
    {
      id: 'mint',
      name: 'Fresh Mint',
      colors: {
        primary: '#00897b',
        secondary: '#e0f2f1',
        background: '#ffffff',
        text: '#000000',
      },
      preview: 'linear-gradient(45deg, #00897b, #e0f2f1)',
    },
  ]

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Choose Theme">
      <div className="py-6">
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
          {themes.map((theme) => (
            <div
              key={theme.id}
              className={`relative rounded-lg overflow-hidden cursor-pointer transition-transform hover:scale-105 ${
                currentThemeId === theme.id ? 'ring-2 ring-primary ring-offset-2' : ''
              }`}
              onClick={() => onSelectTheme(theme.id)}
            >
              <div
                className="h-24 w-full"
                style={{
                  background: theme.preview,
                }}
              />
              <div className="absolute bottom-0 left-0 right-0 bg-black/50 text-white p-2 text-sm text-center">
                {theme.name}
              </div>
              {currentThemeId === theme.id && (
                <div className="absolute top-2 right-2 bg-primary text-white rounded-full p-1">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
                    <path
                      fillRule="evenodd"
                      d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                      clipRule="evenodd"
                    />
                  </svg>
                </div>
              )}
            </div>
          ))}
        </div>
      </div>
    </Modal>
  )
}

export default ThemeModal
