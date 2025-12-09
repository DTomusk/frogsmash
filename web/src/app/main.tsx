import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { HashRouter } from 'react-router-dom'
import Background from './providers/Background.tsx'
import { AppThemeProvider } from './providers/ThemeProvider.tsx'
import { TenantProvider } from './providers/TenantProvider.tsx'

const queryClient = new QueryClient()

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <HashRouter>
        <TenantProvider>
          <AppThemeProvider>
            <Background />
            <App />
          </AppThemeProvider>
        </TenantProvider>
      </HashRouter>
    </QueryClientProvider>
  </StrictMode>,
)
