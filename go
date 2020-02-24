#!/bin/bash

## 파라미터가 없으면 종료 
if [ "$#" -lt 2 ]; then
    echo "$# is Illegal number of parameters."
    echo "Usage: $0 [options]"
    exit 1
fi


## 0. open_manipulator
if [ "$1" == "0" ]; then
  if [ "$2" == "1" ]; then
    echo "roslaunch open_manipulator_gazebo open_manipulator_gazebo.launch"
    `roslaunch open_manipulator_gazebo open_manipulator_gazebo.launch`
  elif [ "$2" == "2" ]; then
    if [ "$3" == "true" ]; then
      echo "roslaunch open_manipulator_controller open_manipulator_controller.launch use_platform:=true"
      `roslaunch open_manipulator_controller open_manipulator_controller.launch use_platform:=true`
    else
      echo "roslaunch open_manipulator_controller open_manipulator_controller.launch use_platform:=false"
      `roslaunch open_manipulator_controller open_manipulator_controller.launch use_platform:=false`
    fi 
  elif [ "$2" == "3" ]; then
    echo "roslaunch open_manipulator_control_gui open_manipulator_control_gui.launch"
    `roslaunch open_manipulator_control_gui open_manipulator_control_gui.launch`
  fi



## robot1_controller
elif [ "$1" == "1" ]; then
  if [ "$2" == "1" ]; then
    echo "roslaunch open_manipulator_gazebo open_manipulator_gazebo.launch use_robot_name:=robot1_controller"
    `roslaunch open_manipulator_gazebo open_manipulator_gazebo.launch use_robot_name:=robot1_controller`
  elif [ "$2" == "2" ]; then
    if [ "$3" == "true" ]; then
      echo "roslaunch open_manipulator_controller open_manipulator_controller.launch use_platform:=true use_robot_name:=robot1_controller dynamixel_usb_port:=/dev/ttyUSB1"
      `roslaunch open_manipulator_controller open_manipulator_controller.launch use_platform:=true use_robot_name:=robot1_controller dynamixel_usb_port:=/dev/ttyUSB1`
    else 
      echo "roslaunch open_manipulator_controller open_manipulator_controller.launch use_platform:=false use_robot_name:=robot1_controller"
      `roslaunch open_manipulator_controller open_manipulator_controller.launch use_platform:=false use_robot_name:=robot1_controller`
    fi
  elif [ "$2" == "3" ]; then
    echo "roslaunch open_manipulator_control_gui open_manipulator_control_gui.launch robot_name:=robot1_controller"
    `roslaunch open_manipulator_control_gui open_manipulator_control_gui.launch robot_name:=robot1_controller`
  fi



## robot2_controller
elif [ "$1" == "2" ]; then
  if [ "$2" == "1" ]; then
    echo "roslaunch open_manipulator_gazebo open_manipulator_gazebo.launch use_robot_name:=robot2_controller"
    `roslaunch open_manipulator_gazebo open_manipulator_gazebo.launch use_robot_name:=robot2_controller`
  elif [ "$2" == "2" ]; then
    if [ "$3" == "true" ]; then
      echo "roslaunch open_manipulator_controller open_manipulator_controller.launch use_platform:=true use_robot_name:=robot2_controller"
      `roslaunch open_manipulator_controller open_manipulator_controller.launch use_platform:=true use_robot_name:=robot2_controller`
    else
      echo "roslaunch open_manipulator_controller open_manipulator_controller.launch use_platform:=false use_robot_name:=robot2_controller"
      `roslaunch open_manipulator_controller open_manipulator_controller.launch use_platform:=false use_robot_name:=robot2_controller`
    fi 
  elif [ "$2" == "3" ]; then
    echo "roslaunch open_manipulator_control_gui open_manipulator_control_gui.launch robot_name:=robot2_controller"
    #`roslaunch open_manipulator_control_gui open_manipulator_control_gui.launch robot_name:=robot2_controller`
  fi

fi

 
