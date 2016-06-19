'use strict';

angular.module('newshoundApp')
    .factory('config', ['$location',
        function($location) {
            return {
                apiHost: function() {
                    var apiHost = "svc/newshound-api/v1";
                    if ($location.host().indexOf('newsai.org') == -1) {
//                        apiHost = "http://137.116.116.97:8080/"+ apiHost;
                        apiHost = "http://newshound1.newsai.org:8080/" + apiHost;
                    }
                    return apiHost;
                }
            };
        }
    ]);
