require 'rails_helper'

RSpec.describe "flags/index", type: :view do
  before(:each) do
    assign(:flags, [
      Flag.create!(
        :name => "Name",
        :enabled => false
      ),
      Flag.create!(
        :name => "Name",
        :enabled => false
      )
    ])
  end

  it "renders a list of flags" do
    render
    assert_select "tr>td", :text => "Name".to_s, :count => 2
    assert_select "tr>td", :text => false.to_s, :count => 2
  end
end
