## Nagios check-ram Plugin

### Installation 

To execute seperatly the plugin, get the source code and build it in your $GOPATH :

    $ go get github.com/aydintd/nagios-check-ram
    $ go install
    $ nagios-check-ram -w 20 -c 25                                                                                                                                                                                              
    
    Memory: WARNING - Total: 8085380 KB - Used: 4830928 KB - Memory Usage: %20|TOTAL=8085380;;;; USED=1633768;;;; FREE=4830928;;;; CACHE=1620684;;;; 

### LICENSE      

	nagios-ram-check - Nagios RAM check plugin
    Copyright (C) 2016  Aydin Doyak <aydintd@gmail.com>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.


