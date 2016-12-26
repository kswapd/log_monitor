
		myjson :=`{
		  "type": "container",
		  "data": {
		    "container_uuid": "b7a37421-647e-4821-8686-aadfff162e14",
		    "environment_id": "123",
		    "namespace": "name123",
		    "container_name": "lonely_blackwell",
		    "timestamp": "2016-12-12T05:26:00.759+00:00",
		    "log_info": {
		      "log_time": "2016-12-12T05:25:57.000+00:00",
		      "source": "stdout",
		      "message": "2016/12/12 05:25:57 ### End of ScrapeGlobalStatus.\r"
		    }
		  }
		}`


	json.Unmarshal([]byte(myjson), &containerJson)
   // fmt.Println(containerJson)

        mjson := `{
      "timestamp": "2016-12-13T06:05:55.782972713Z",
      "container_uuid": "",
      "environment_id": "1234",
      "container_name": "",
      "namespace": "",
      "stats": [
        {
          "timestamp": "2016-12-13T06:05:55.57493879Z",
          "container_cpu_usage_seconds_total": 11716012449561,
          "container_cpu_user_seconds_total": 6815830000000,
          "container_cpu_system_seconds_total": 4464280000000,
          "container_memory_usage_bytes": 44470272,
          "container_memory_limit_bytes": 18446744073709551615,
          "container_memory_cache": 32768,
          "container_memory_rss": 44437504,
          "container_memory_swap": 0,
          "container_network_receive_bytes_total": 51616316,
          "container_network_receive_packets_total": 554902,
          "container_network_receive_packets_dropped_total": 0,
          "container_network_receive_errors_total": 0,
          "container_network_transmit_bytes_total": 2330468437,
          "container_network_transmit_packets_total": 1327891,
          "container_network_transmit_packets_dropped_total": 0,
          "container_network_transmit_errors_total": 0,
          "container_filesystem": [
            {
              "container_filesystem_name": "/dev/vda1",
              "container_filesystem_type": "vfs",
              "container_filesystem_capacity": 84516978688,
              "container_filesystem_usage": 3203698688
            }
          ],
          "container_diskio_service_bytes_async": 0,
          "container_diskio_service_bytes_read": 0,
          "container_diskio_service_bytes_sync": 0,
          "container_diskio_service_bytes_total": 0,
          "container_diskio_service_bytes_write": 0,
          "container_tasks_state_nr_sleeping": 0,
          "container_tasks_state_nr_running": 0,
          "container_tasks_state_nr_stopped": 0,
          "container_tasks_state_nr_uninterruptible": 0,
          "container_tasks_state_nr_io_wait": 0
        }
      ]
    }` 
    mystats := `

    [
        {
          "timestamp": "2016-12-13T06:05:55.57493879Z",
          "container_cpu_usage_seconds_total": 11716012449561,
          "container_cpu_user_seconds_total": 6815830000000,
          "container_cpu_system_seconds_total": 4464280000000,
          "container_memory_usage_bytes": 44470272,
          "container_memory_limit_bytes": 18446744073709551615,
          "container_memory_cache": 32768,
          "container_memory_rss": 44437504,
          "container_memory_swap": 0,
          "container_network_receive_bytes_total": 51616316,
          "container_network_receive_packets_total": 554902,
          "container_network_receive_packets_dropped_total": 0,
          "container_network_receive_errors_total": 0,
          "container_network_transmit_bytes_total": 2330468437,
          "container_network_transmit_packets_total": 1327891,
          "container_network_transmit_packets_dropped_total": 0,
          "container_network_transmit_errors_total": 0,
          "container_filesystem": [
            {
              "container_filesystem_name": "/dev/vda1",
              "container_filesystem_type": "vfs",
              "container_filesystem_capacity": 84516978688,
              "container_filesystem_usage": 3203698688
            }
          ],
          "container_diskio_service_bytes_async": 0,
          "container_diskio_service_bytes_read": 0,
          "container_diskio_service_bytes_sync": 0,
          "container_diskio_service_bytes_total": 0,
          "container_diskio_service_bytes_write": 0,
          "container_tasks_state_nr_sleeping": 0,
          "container_tasks_state_nr_running": 0,
          "container_tasks_state_nr_stopped": 0,
          "container_tasks_state_nr_uninterruptible": 0,
          "container_tasks_state_nr_io_wait": 0
        },
		{
          "timestamp": "2016-12-13T06:05:65.57493879Z",
          "container_cpu_usage_seconds_total": 11716012449561,
          "container_cpu_user_seconds_total": 6815830000000,
          "container_cpu_system_seconds_total": 4464280000000,
          "container_memory_usage_bytes": 44470272,
          "container_memory_limit_bytes": 18446744073709551615,
          "container_memory_cache": 32768,
          "container_memory_rss": 44437504,
          "container_memory_swap": 0,
          "container_network_receive_bytes_total": 51616316,
          "container_network_receive_packets_total": 554902,
          "container_network_receive_packets_dropped_total": 0,
          "container_network_receive_errors_total": 0,
          "container_network_transmit_bytes_total": 2330468437,
          "container_network_transmit_packets_total": 1327891,
          "container_network_transmit_packets_dropped_total": 0,
          "container_network_transmit_errors_total": 0,
          "container_filesystem": [
            {
              "container_filesystem_name": "/dev/vda1",
              "container_filesystem_type": "vfs",
              "container_filesystem_capacity": 84516978688,
              "container_filesystem_usage": 3203698688
            }
          ],
          "container_diskio_service_bytes_async": 0,
          "container_diskio_service_bytes_read": 0,
          "container_diskio_service_bytes_sync": 0,
          "container_diskio_service_bytes_total": 0,
          "container_diskio_service_bytes_write": 0,
          "container_tasks_state_nr_sleeping": 0,
          "container_tasks_state_nr_running": 0,
          "container_tasks_state_nr_stopped": 0,
          "container_tasks_state_nr_uninterruptible": 0,
          "container_tasks_state_nr_io_wait": 0
        }

      ]

    `

