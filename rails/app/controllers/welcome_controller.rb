# frozen_string_literal: true

class WelcomeController < ApplicationController
  def index
    respond_to do |format|
      format.html
    end
  end
end
