function EventsCtrl($scope, $http) {
  // The list of events in display.
  $scope.events = [];
  // The fields in the event creation dialog.
  $scope.newEvent = {};

  // Display an error using an alert dialog.
  var alertError = function (data, status) {
    alert("code " + status + ": " + data);
  };

  // Fetches all the events from the API.
  var fetchEvents = function () {
    return $http
      .get("/api/events")
      .error(alertError)
      .success(function (data) {
        $scope.events = data;
      });
  };

  // Adds a new event throught the API.
  $scope.addEvent = function () {
    $http
      .post("/api/events", $scope.newEvent)
      .error(alertError)
      .success(function () {
        fetchEvents().then(function () {
          // If everything worked, clear the dialog.
          $scope.event = {};
          // Fetch again after a bit, in case of eventual consistency.
          setTimeout(fetchEvents, 1000);
        });
      });
  };

  // Fetch the list of events from the API.
  fetchEvents();
}
