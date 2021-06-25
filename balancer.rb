#####################################################################
#   								    #
#           Ciprand reverse proxy and load balancer	            #
# 								    #
#	    this proxy server chooses a random server for           #
#           each request to handle the incoming traffic 	    # 
#                                                                   #
#####################################################################

require 'cgi'
require 'net/http'
require 'socket'


$servers = [
	ENV["SERVER1"],
	ENV["SERVER2"], #
	ENV["SERVER3"], # servers subdomain
	ENV["SERVER4"],	#
	ENV["SERVER5"]	
]

$baseurl = ENV["BASEURL"] # main domain
server = TCPServer.new(ENV["PORT"]) # Port : proxy port

def ServerGet(request)
	begin
		uri = URI(request)
		params = CGI::parse(uri.query)
		server = $servers[rand($servers.length)]
		server = "https://#{server}.#{$baseurl}/api?len=#{params['len'][0]}&count=#{params['count'][0]}"
		resp = Net::HTTP.get_response(URI(server))
		b = resp.body
	rescue
		b = false
	end
	return b

end

loop do 
	client = server.accept 
	Thread.start {
		request = client.readpartial(3000)

		if request.split[0] != "GET" 
			head = "HTTP/1.1 405\r\n"
			head += "Content-Length: 18\r\n"
			head += "Content-Type: plain/text\r\n"
			head += "\r\n"
			head += "Method not allowed\r\n"
			client.write(head)
			client.close
			Thread.exit()
		end
		res = ServerGet(request.split[1])
		code = 200
		if !res
			res = '{"status": "error"}'
			code = 500
		end

		head = "HTTP/1.1 #{code}\r\n"
		head += "Content-Length: #{res.length}\r\n"
		head += "Content-Type: application/json\r\n"
		head += "\r\n"
		head += "#{res}\r\n"
		client.write(head)
		client.close
	}
end



