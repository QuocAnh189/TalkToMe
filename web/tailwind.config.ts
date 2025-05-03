import type { Config } from 'tailwindcss'

const config: Config = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {
      maxWidth: {
        container: '1440px',
        contentContainer: '1140px',
        containerSmall: '1024px',
        containerXs: '768px',
      },

      screens: {
        xs: '320px',
        sm: '375px',
        sml: '500px',
        md: '667px',
        mdl: '768px',
        lg: '960px',
        lgl: '1024px',
        xl: '1280px',
      },
    },
    colors: {
      primary: '#3490dc', 
      secondary: '#6cb2eb', 
      accent: '#f6993f', 
      neutral: '#3d4451', 
      'base-100': '#ffffff', 
      info: '#2094f3',
      success: '#38c172',
      warning: '#ffed4a',
      error: '#e3342f',
      'message-sent': '#e2e8f0',
      'message-text-sent': '#2d3748', 
      'message-received': '#ffffff', 
      'message-text-received': '#2d3748', 
      'sidebar-bg': '#f7fafc', 
      'header-bg': '#ffffff',
      'border-color': '#e5e7eb', 
    },
  },
  plugins: [],
}
export default config
