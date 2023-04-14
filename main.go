package main

/*
#cgo LDFLAGS: -lvulkan
#include <stdlib.h>
#include <assert.h>
#include "vulkan.h"

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

struct handles {
	VkInstance instance;
	VkDevice device;
};
typedef struct handles handles;

handles *start()  {
	handles *handles = malloc(sizeof( handles));

	PFN_vkCreateInstance createInstance = (PFN_vkCreateInstance)(vkGetInstanceProcAddr(NULL, "vkCreateInstance"));

	VkInstanceCreateInfo *createInfo = createInstanceCreateInfo();

	VkResult res = createInstance(createInfo, NULL, &handles->instance);
	assert(res == VK_SUCCESS);

	free(createInfo->pApplicationInfo);
	free(createInfo);

	PFN_vkEnumeratePhysicalDevices enumeratePhysicalDevices = (PFN_vkEnumeratePhysicalDevices)(vkGetInstanceProcAddr(handles->instance, "vkEnumeratePhysicalDevices"));
	PFN_vkCreateDevice createDevice = (PFN_vkCreateDevice)(vkGetInstanceProcAddr(handles->instance, "vkCreateDevice"));

	uint32_t count;

	res = enumeratePhysicalDevices(handles->instance, &count, NULL);
	assert(res == VK_SUCCESS);

	VkPhysicalDevice *physicalDeviceHandles = malloc(count*sizeof(VkPhysicalDevice));
	res = enumeratePhysicalDevices(handles->instance, &count, physicalDeviceHandles);
	assert(res == VK_SUCCESS);

	VkDeviceCreateInfo *createDeviceInfo = createDeviceCreateInfo();
	res = createDevice(physicalDeviceHandles[0], createDeviceInfo, NULL, &handles->device);
	assert(res == VK_SUCCESS);

	free(physicalDeviceHandles);
	free(createDeviceInfo->pQueueCreateInfos);
	free(createDeviceInfo);

	return handles;
}

void end(handles *handles) {
	PFN_vkDestroyInstance destroyInstance = (PFN_vkDestroyInstance)(vkGetInstanceProcAddr(handles->instance, "vkDestroyInstance"));
	PFN_vkDestroyDevice destroyDevice = (PFN_vkDestroyDevice)(vkGetInstanceProcAddr(handles->instance, "vkDestroyDevice"));

	destroyDevice(handles->device, NULL);
	destroyInstance(handles->instance, NULL);
	free(handles);
}

*/
import "C"
import (
	"fmt"
	"runtime"
)

func main() {

	for i := 0; i < 10000; i++ {
		handles := C.start()
		runtime.GC()
		C.end(handles)
		fmt.Println(i)
	}
}
