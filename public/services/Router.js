// Creating Router as a singelton
const Router = {
  // Actions to take when a Router is created
  init: () => {
    // User navigates back to previous page. We want delete that previous state from the history, so he would not be able to go back to that page (because it doesn't make sense)
    window.addEventListener('popstate', () => {
      Router.go(location.pathname, false);
    });

    // Go to the initial route. the search is optional and is URL query and will be passed after the '?' mark
    Router.go(location.pathname + location.search);
  },

  // addToHistory will enable to control the option of the user to go back to previous page
  go: (route, addToHistory = true) => {
    if (addToHistory) {
      // User navigates to a page. We want save the previous state in the history, so he would be able to go back if he wants
      history.pushState(null, '', route);
    }
  },
};
