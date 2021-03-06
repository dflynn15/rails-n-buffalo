require 'rails_helper'

RSpec.describe "flags/show", type: :view do
  before(:each) do
    @flag = assign(:flag, Flag.create!(
      :name => "Name",
      :enabled => false
    ))
  end

  it "renders attributes in <p>" do
    render
    expect(rendered).to match(/Name/)
    expect(rendered).to match(/false/)
  end
end
