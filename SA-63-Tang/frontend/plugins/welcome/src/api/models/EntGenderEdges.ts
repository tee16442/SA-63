/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Patient
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntCustomer,
    EntCustomerFromJSON,
    EntCustomerFromJSONTyped,
    EntCustomerToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntGenderEdges
 */
export interface EntGenderEdges {
    /**
     * Customer2 holds the value of the customer2 edge.
     * @type {Array<EntCustomer>}
     * @memberof EntGenderEdges
     */
    customer2?: Array<EntCustomer>;
}

export function EntGenderEdgesFromJSON(json: any): EntGenderEdges {
    return EntGenderEdgesFromJSONTyped(json, false);
}

export function EntGenderEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntGenderEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'customer2': !exists(json, 'customer2') ? undefined : ((json['customer2'] as Array<any>).map(EntCustomerFromJSON)),
    };
}

export function EntGenderEdgesToJSON(value?: EntGenderEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'customer2': value.customer2 === undefined ? undefined : ((value.customer2 as Array<any>).map(EntCustomerToJSON)),
    };
}


