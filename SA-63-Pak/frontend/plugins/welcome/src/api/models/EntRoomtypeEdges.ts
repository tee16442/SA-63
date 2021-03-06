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
 * @interface EntRoomtypeEdges
 */
export interface EntRoomtypeEdges {
    /**
     * Room1 holds the value of the Room1 edge.
     * @type {Array<EntRoom>}
     * @memberof EntRoomtypeEdges
     */
    room1?: Array<EntRoom>;
}

export function EntRoomtypeEdgesFromJSON(json: any): EntRoomtypeEdges {
    return EntRoomtypeEdgesFromJSONTyped(json, false);
}

export function EntRoomtypeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRoomtypeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'room1': !exists(json, 'room1') ? undefined : ((json['room1'] as Array<any>).map(EntRoomFromJSON)),
    };
}

export function EntRoomtypeEdgesToJSON(value?: EntRoomtypeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'room1': value.room1 === undefined ? undefined : ((value.room1 as Array<any>).map(EntRoomToJSON)),
    };
}


