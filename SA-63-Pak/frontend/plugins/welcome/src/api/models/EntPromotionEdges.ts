/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
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
    EntRoom,
    EntRoomFromJSON,
    EntRoomFromJSONTyped,
    EntRoomToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPromotionEdges
 */
export interface EntPromotionEdges {
    /**
     * Room2 holds the value of the Room2 edge.
     * @type {Array<EntRoom>}
     * @memberof EntPromotionEdges
     */
    room2?: Array<EntRoom>;
}

export function EntPromotionEdgesFromJSON(json: any): EntPromotionEdges {
    return EntPromotionEdgesFromJSONTyped(json, false);
}

export function EntPromotionEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPromotionEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'room2': !exists(json, 'room2') ? undefined : ((json['room2'] as Array<any>).map(EntRoomFromJSON)),
    };
}

export function EntPromotionEdgesToJSON(value?: EntPromotionEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'room2': value.room2 === undefined ? undefined : ((value.room2 as Array<any>).map(EntRoomToJSON)),
    };
}


