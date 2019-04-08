Rails.application.routes.draw do
  resources :projects do
    resources :flags
  end
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  root 'welcome#index'
end
