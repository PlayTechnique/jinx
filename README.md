Jinx: a jinkies coordination tool

# Quick Usage:
For more thorough documentation, read the source or check out the --help flags for each subcommand.

`jinkies serve start` - starts a jinkies on localhost 😯

`jinkies serve start -o hostconfig.yml` - the `-o` flag allows the end user to specify any of the docker engine [hostconfig options](https://pkg.go.dev/github.com/docker/docker@v20.10.12+incompatible/api/types/container#HostConfig) 😯

`jinkies serve start -c containerconfig.yml` - the `-c` flag allows the end user to specify any of the docker engine [containerconfig options](https://pkg.go.dev/github.com/docker/docker@v20.10.12+incompatible/api/types/container#Config) 😯

Examples of how these yml files are structured are in the examples/ directory.

## Global Config File
For custom jinx options, we support a config file called jinx.yml. It must be present in the same directory as the jinx binary is invoked from.
The flags this supports are documented in the examples/ directory.

## Developers
Check out [Developer.md](DEVELOPER.md) for hints and tricks for getting started. We want your PR!



                     *╗╖                                                        
                        ╫▒╣╦                           ╓#≥#▓▀Γ`                 
                        ]▒░░▒                       ╓▒░░▒╝`                     
                        ╟▒░░░b                     ╬░░░╠╣                       
                       ]╬▄▓▓██                   ▄██▓▓▄╠▓                       
                       ███╬╬██▄▒             ,╦╣███╬╬████                       
                      ████╩╚░░░░▒╓       ╓≤▒╠░░▒░░╠╟█████                       
                      ███░░░░░░░░░░▒╚░░░░░░░░░░░░░░▒╬███▌                       
                      █▓▒░░░░░░░░░░░░░░░░░░░░░░░░▒╬╬╬╬██                        
                      ╙╬╬▒░░╟▒╠╩▒░░▒░░░░░░▒▒░░▒Å╣╬╬╬╬╬█─                        
                       ╣╬╬▒▒▓▄╙╕│╙╙╙╝╙╙▒╙╙║▓╬╬╬╬╬╬╬╬╬╩                          
                        ╙╣╬▓▌▒╩░░░░░#╠╠╚─╠╦▓╣╬╠╟╬╬╬╩,                           
                          ╙▒╙╙░│░░░░,║╝╝▀│╠░░╚▒║╬╣▀                             
                           Ü¡¡░╬░░░░░││¡░░░╟╗╣╝╨`                               
                            Q░░░░░░░░░░░░╝`                                     
                             "░∩5░░░░░▒╜                                        
                                \░≡Å╙╚µ                                         
                                   ██▓█                                         
                                  ⌠╠╬╙░▒,                                       
                               ,≤░▐╬▓░│╚░▒▄                                     
                             #▓╬▄░░│░░▄#▓███▓µ                                  
                            ║╬╬╬╬╬╬▓▓╬╬╬╬╬╬▓██▌                                 
                           ]▓╬╬╬╬╬╣█╬╬╬╬╬╬╬╬╬██▌                                
                           ▓╬╬╬╬╬╣███╬╬╬╬╬╬╬╬╬╣██                               
                         ╓▓╬╬╬╬╬╬████╬╬╬╬╬╬╬╬╣▀╙▀▀═                             
                            ▀██▓███████╣█Γ╟███                                  
                             ██╣███████╫▌ j███                                  
                             ╫██╟█████▓█   ██▌                                  
                             ╟██j█████╫▌   ██▌                                  
                             █████████▓╫█  ╟██                                  
                            ▐██████████▓╣█ ▐███                                 
                            ████████████╬██▐████                                
                            ████████████╬█████╬█▌                               
                            █████████████╫████╬██                               
                            █████████████╬█████╫█                               
                            █████████████╬╣████╬█µ                              
                           ]██████████████╬████╬█▌                              
                           ▐██╣███████████╬╬████╬█                              
                           ╟██╣████████████╬████╬█▌                             
                           ███╣██████▀████╬█████╬██                             
                           ███╬██╬╬╣█  ╬╬╬╬█████╬╬█▌                            
                          @j▒▀╩▀█╬╬╬█⌐ ╟╬╬╬▓▀█▀█▀╠╙                             
                          ▒╣▓║╬╖,╟╬╬▓b  ▓╬╬█▓▐▐░▒▒b                             
                          └╟╚╠╙╙⌐ ███▌  ╟████▐▐]╩╘Γ                             
                            ⌐ "   ╫██▌   ████▒▌╨                                
                                  ╘██▌   ████                                   
                                   ╬╣▌   └╬╬█                                   
                                   ╬╣▌    ╬╬█                                   
                                   ▓▓█    ▓╬▓▌                                  
                                   ▓╬█▌   ╬╬╬█▄                                 
                                   ████   █████                                 
                                   ████   ╟████                                 
                                   ████    ████                                 
                                   ███Γ    ╟███                                 
                                   █╬█      ╬╣█                                 
                                   ▓╣▌      ▓╣█                                 
                                   ▓╣▌      ╟╬█                                 
                                  ╟█▒█      █╣█▌                                
                                  ████      ╟╣█▌                                
                                 ╫██▓█      ╟╣██                                
                                ▄███▒█▌     ▐▒██                                
                               ▓███▒▓███    █╫██▌                               
                            ,▓████▓█████    █████▄                              
                        ,▄▓████╬▓███████    █╬████▌                             
                     ▄▓██╬╬╬██████╬╬████    ╟███████▄                           
                    Φ████████████╬╬╬████⌐   ▐█╬███████▌                         
                   ╣╬╬╬╬╬╬╬╬╬╬╣██╝╬╣████¬   j██████████▌                        
                  ]╬╬╬╬╬╬╬╣╬╬▓██⌐   └        ███████████▄                       
                                             ╬╬╬╬╬╬╬╬╬╬▓█▌                      
                                             ╬╬╬╬╬╬╬╬╬╬╣██                      
                                              ╙╝╣╬╬╬╬╬╬▓▓▀"                     
