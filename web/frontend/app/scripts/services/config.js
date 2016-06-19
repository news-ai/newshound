'use strict';

angular.module('newshoundApp')
    .factory('config', ['$location',
        function($location) {
            return {
                apiHost: function() {
                    return "http://newshound1.newsai.org:8080/svc/newshound-api/v1";
                }
            };
        }
    ]);
