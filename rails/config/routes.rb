Rails.application.routes.draw do
  devise_for :users, controllers: { omniauth_callbacks: 'users/omniauth_callbacks' }

  # API
  namespace :api, defaults: { format: :json } do
    namespace :v1 do
      resources :projects, only: [:show] do
        resources :flags, only: [:show]
      end
      get '/me' => "credentials#me"
    end
  end

  # UI
  resources :projects do
    resources :flags
  end

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  root 'welcome#index'
end
