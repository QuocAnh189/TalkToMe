import type { Config } from 'tailwindcss';

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
      colors: {
        'app-primary': '#3490dc',
        'app-secondary': '#6cb2eb',
        'app-secondaryLight': '#e3f2fd',
        'app-accent': '#f6993f',
        'app-neutral': '#3d4451',
        'app-base-100': '#ffffff',
        'app-info': '#2094f3',
        'app-success': '#38c172',
        'app-warning': '#ffed4a',
        'app-error': '#e3342f',
        'app-message-sent': '#e2e8f0',
        'app-message-text-sent': '#2d3748',
        'app-message-received': '#ffffff',
        'app-message-text-received': '#2d3748',
        'app-sidebar-bg': '#f7fafc',
        'app-header-bg': '#ffffff',
        'app-border-color': '#e5e7eb',
      },
    },
  },
  plugins: [],
};
export default config;