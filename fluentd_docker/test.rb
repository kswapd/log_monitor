require 'json';require "open-uri";uri = 'http://169.254.169.250/2015-12-19/containers/';html_response = nil;uuid = nil;env_id = nil;open(uri,"Accept" => "application/json") do |http| html_response = http.read end; container_array = JSON.parse(html_response); container_array.each {|info| puts info['uuid']; if info['external_id']=='dc4a6e148c2daf322523aafc8827a34eebe2e7f4b3481074044628d9bb14296d'; uuid=info['uuid'];env_id=info['labels']['io.rancher.container.name']; end; }
    puts uuid
    puts env_id

