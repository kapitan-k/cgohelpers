#include "helpers.h"


#include <stdlib.h>




void free_multi(void **datas, size_t num_datas) {
	for (size_t i = 0; i < num_datas; i++) {
		if (datas[i] != NULL) {
			free(datas[i]); 
		}
	}
}