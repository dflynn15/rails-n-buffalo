Rails.application.routes.draw do
  # api
  namespace :api do
    namespace :v1 do
      resources :profiles
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
