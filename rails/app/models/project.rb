class Project < ApplicationRecord
  belongs_to :user

  has_many :flags, dependent: :destroy
  validates_presence_of :name
end
