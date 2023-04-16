package main

/*
#cgo LDFLAGS: -lvulkan
#include <stdlib.h>
#include <assert.h>
#include <pthread.h>
#include "vulkan.h"

VkInstance instance;
VkDevice device;

VkInstanceCreateInfo *createInstanceCreateInfo()  {
	VkInstanceCreateInfo *createInfo = malloc(sizeof(VkInstanceCreateInfo));
	VkApplicationInfo *appInfo = malloc(sizeof(VkApplicationInfo));

	appInfo->sType = VK_STRUCTURE_TYPE_APPLICATION_INFO;
	appInfo->pNext = NULL;
	appInfo->pApplicationName = NULL;
	appInfo->pEngineName = NULL;
	appInfo->applicationVersion = 0;
	appInfo->engineVersion = 0;
	appInfo->apiVersion = VK_API_VERSION_1_0;

	createInfo->sType = VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO;
	createInfo->flags = 0;
	createInfo->pNext = NULL;
	createInfo->pApplicationInfo = appInfo;
	createInfo->enabledExtensionCount = 0;
	createInfo->ppEnabledExtensionNames = NULL;
	createInfo->enabledLayerCount = 0;
	createInfo->ppEnabledLayerNames = NULL;

	return createInfo;
}

VkDeviceQueueCreateInfo *createDeviceQueueCreateInfo() {
	VkDeviceQueueCreateInfo *createInfo = malloc(sizeof(VkDeviceQueueCreateInfo));

	float *prioritiesPtr = malloc(sizeof(float));
	*prioritiesPtr = 0;

	createInfo->sType = VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO;
	createInfo->flags = 0;
	createInfo->pNext = NULL;
	createInfo->queueCount = 1;
	createInfo->queueFamilyIndex = 0;
	createInfo->pQueuePriorities = prioritiesPtr;

	return createInfo;
}

VkDeviceCreateInfo *createDeviceCreateInfo() {
	VkDeviceCreateInfo *createInfo = malloc(sizeof(VkDeviceCreateInfo));
	// Alloc queue families
	VkDeviceQueueCreateInfo *queueFamilyPtr = createDeviceQueueCreateInfo();

	createInfo->sType = VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO;
	createInfo->flags = 0;
	createInfo->pNext = NULL;
	createInfo->queueCreateInfoCount = 1;
	createInfo->pQueueCreateInfos = queueFamilyPtr;
	createInfo->enabledLayerCount = 0;
	createInfo->ppEnabledLayerNames = NULL;
	createInfo->enabledExtensionCount = 0;
	createInfo->ppEnabledExtensionNames = NULL;
	createInfo->pEnabledFeatures = NULL;

	return createInfo;
}

void start()  {
	PFN_vkCreateInstance createInstance = (PFN_vkCreateInstance)(vkGetInstanceProcAddr(NULL, "vkCreateInstance"));

	VkInstanceCreateInfo *createInfo = createInstanceCreateInfo();

	VkResult res = createInstance(createInfo, NULL, &instance);
	assert(res == VK_SUCCESS);

	free(createInfo->pApplicationInfo);
	free(createInfo);

	PFN_vkEnumeratePhysicalDevices enumeratePhysicalDevices = (PFN_vkEnumeratePhysicalDevices)(vkGetInstanceProcAddr(instance, "vkEnumeratePhysicalDevices"));
	PFN_vkCreateDevice createDevice = (PFN_vkCreateDevice)(vkGetInstanceProcAddr(instance, "vkCreateDevice"));

	uint32_t *count = malloc(sizeof(uint32_t));

	res = enumeratePhysicalDevices(instance, count, NULL);
	assert(res == VK_SUCCESS);

	VkPhysicalDevice *physicalDeviceHandles = malloc((*count)*sizeof(VkPhysicalDevice));
	res = enumeratePhysicalDevices(instance, count, physicalDeviceHandles);
	assert(res == VK_SUCCESS);

	free(count);

	VkDeviceCreateInfo *createDeviceInfo = createDeviceCreateInfo();
	res = createDevice(physicalDeviceHandles[0], createDeviceInfo, NULL, &device);
	assert(res == VK_SUCCESS);

	free(physicalDeviceHandles);
	free(createDeviceInfo->pQueueCreateInfos);
	free(createDeviceInfo);
}

void end() {
	PFN_vkDestroyInstance destroyInstance = (PFN_vkDestroyInstance)(vkGetInstanceProcAddr(instance, "vkDestroyInstance"));
	PFN_vkDestroyDevice destroyDevice = (PFN_vkDestroyDevice)(vkGetInstanceProcAddr(instance, "vkDestroyDevice"));

	destroyDevice(device, NULL);
	destroyInstance(instance, NULL);
}

*/
import "C"
import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"math"
	"runtime/debug"
)

var done = make(chan bool)

func handler() {
	C.start()

	var e errgroup.Group
	for i := 0; i < 5; i++ {
		e.Go(func() error {
			return doWork(1, 1000000)
		})
	}

	last := C.pthread_self()
	_ = e.Wait()
	new := C.pthread_self()
	if new != last {
		fmt.Println("thread", new)
	}

	C.end()

	done <- true
}

func doWork(num1, num2 int) error {
	n := num1
	for n <= num2 {
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				break
			}
		}
		n++
	}

	return nil
}

func main() {
	//procTimeBeginPeriod.Call(uintptr(1))
	debug.SetGCPercent(-1)

	for i := 0; i < 100; i++ {
		go handler()
		<-done
		fmt.Println(i)
	}

}
