package storage

type UserStorage struct {

}

func (self *UserStorage) get (userId, campaignKey int) {
	/*
		To retrieve the stored variation for the user_id and
		campaign_key

		Args:
		    user_id (str): User ID for which data needs to be retrieved.
		    campaign_key (str): Campaign key to identify the campaign for
		    which stored variation should be retrieved.

		Returns:
		    user_data (dict): user-variation mapping
  */
  //No need to use pass or continue 
}

func (self *UserStorage) set (userData int) {
	/*
		To store the the user variation-mapping

		Args:
		    user_data (dict): user-variation mapping
  */
  //No need to use pass or continue 
}
