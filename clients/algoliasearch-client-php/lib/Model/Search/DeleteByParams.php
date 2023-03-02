<?php

// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

namespace Algolia\AlgoliaSearch\Model\Search;

/**
 * DeleteByParams Class Doc Comment
 *
 * @category Class
 * @package Algolia\AlgoliaSearch
 */
class DeleteByParams extends \Algolia\AlgoliaSearch\Model\AbstractModel implements
    ModelInterface,
    \ArrayAccess,
    \JsonSerializable
{
    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @var string[]
     */
    protected static $modelTypes = [
        'facetFilters' => '\Algolia\AlgoliaSearch\Model\Search\FacetFilters',
        'filters' => 'string',
        'numericFilters' => '\Algolia\AlgoliaSearch\Model\Search\NumericFilters',
        'tagFilters' => '\Algolia\AlgoliaSearch\Model\Search\TagFilters',
        'aroundLatLng' => 'string',
        'aroundRadius' => '\Algolia\AlgoliaSearch\Model\Search\AroundRadius',
        'insideBoundingBox' => 'float[]',
        'insidePolygon' => 'float[]',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'facetFilters' => null,
        'filters' => null,
        'numericFilters' => null,
        'tagFilters' => null,
        'aroundLatLng' => null,
        'aroundRadius' => null,
        'insideBoundingBox' => 'double',
        'insidePolygon' => 'double',
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'facetFilters' => 'facetFilters',
        'filters' => 'filters',
        'numericFilters' => 'numericFilters',
        'tagFilters' => 'tagFilters',
        'aroundLatLng' => 'aroundLatLng',
        'aroundRadius' => 'aroundRadius',
        'insideBoundingBox' => 'insideBoundingBox',
        'insidePolygon' => 'insidePolygon',
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function modelTypes()
    {
        return self::$modelTypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function modelFormats()
    {
        return self::$modelFormats;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'facetFilters' => 'setFacetFilters',
        'filters' => 'setFilters',
        'numericFilters' => 'setNumericFilters',
        'tagFilters' => 'setTagFilters',
        'aroundLatLng' => 'setAroundLatLng',
        'aroundRadius' => 'setAroundRadius',
        'insideBoundingBox' => 'setInsideBoundingBox',
        'insidePolygon' => 'setInsidePolygon',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'facetFilters' => 'getFacetFilters',
        'filters' => 'getFilters',
        'numericFilters' => 'getNumericFilters',
        'tagFilters' => 'getTagFilters',
        'aroundLatLng' => 'getAroundLatLng',
        'aroundRadius' => 'getAroundRadius',
        'insideBoundingBox' => 'getInsideBoundingBox',
        'insidePolygon' => 'getInsidePolygon',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     */
    public function __construct(array $data = null)
    {
        if (isset($data['facetFilters'])) {
            $this->container['facetFilters'] = $data['facetFilters'];
        }
        if (isset($data['filters'])) {
            $this->container['filters'] = $data['filters'];
        }
        if (isset($data['numericFilters'])) {
            $this->container['numericFilters'] = $data['numericFilters'];
        }
        if (isset($data['tagFilters'])) {
            $this->container['tagFilters'] = $data['tagFilters'];
        }
        if (isset($data['aroundLatLng'])) {
            $this->container['aroundLatLng'] = $data['aroundLatLng'];
        }
        if (isset($data['aroundRadius'])) {
            $this->container['aroundRadius'] = $data['aroundRadius'];
        }
        if (isset($data['insideBoundingBox'])) {
            $this->container['insideBoundingBox'] = $data['insideBoundingBox'];
        }
        if (isset($data['insidePolygon'])) {
            $this->container['insidePolygon'] = $data['insidePolygon'];
        }
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }

    /**
     * Gets facetFilters
     *
     * @return \Algolia\AlgoliaSearch\Model\Search\FacetFilters|null
     */
    public function getFacetFilters()
    {
        return $this->container['facetFilters'] ?? null;
    }

    /**
     * Sets facetFilters
     *
     * @param \Algolia\AlgoliaSearch\Model\Search\FacetFilters|null $facetFilters facetFilters
     *
     * @return self
     */
    public function setFacetFilters($facetFilters)
    {
        $this->container['facetFilters'] = $facetFilters;

        return $this;
    }

    /**
     * Gets filters
     *
     * @return string|null
     */
    public function getFilters()
    {
        return $this->container['filters'] ?? null;
    }

    /**
     * Sets filters
     *
     * @param string|null $filters filter the query with numeric, facet and/or tag filters
     *
     * @return self
     */
    public function setFilters($filters)
    {
        $this->container['filters'] = $filters;

        return $this;
    }

    /**
     * Gets numericFilters
     *
     * @return \Algolia\AlgoliaSearch\Model\Search\NumericFilters|null
     */
    public function getNumericFilters()
    {
        return $this->container['numericFilters'] ?? null;
    }

    /**
     * Sets numericFilters
     *
     * @param \Algolia\AlgoliaSearch\Model\Search\NumericFilters|null $numericFilters numericFilters
     *
     * @return self
     */
    public function setNumericFilters($numericFilters)
    {
        $this->container['numericFilters'] = $numericFilters;

        return $this;
    }

    /**
     * Gets tagFilters
     *
     * @return \Algolia\AlgoliaSearch\Model\Search\TagFilters|null
     */
    public function getTagFilters()
    {
        return $this->container['tagFilters'] ?? null;
    }

    /**
     * Sets tagFilters
     *
     * @param \Algolia\AlgoliaSearch\Model\Search\TagFilters|null $tagFilters tagFilters
     *
     * @return self
     */
    public function setTagFilters($tagFilters)
    {
        $this->container['tagFilters'] = $tagFilters;

        return $this;
    }

    /**
     * Gets aroundLatLng
     *
     * @return string|null
     */
    public function getAroundLatLng()
    {
        return $this->container['aroundLatLng'] ?? null;
    }

    /**
     * Sets aroundLatLng
     *
     * @param string|null $aroundLatLng search for entries around a central geolocation, enabling a geo search within a circular area
     *
     * @return self
     */
    public function setAroundLatLng($aroundLatLng)
    {
        $this->container['aroundLatLng'] = $aroundLatLng;

        return $this;
    }

    /**
     * Gets aroundRadius
     *
     * @return \Algolia\AlgoliaSearch\Model\Search\AroundRadius|null
     */
    public function getAroundRadius()
    {
        return $this->container['aroundRadius'] ?? null;
    }

    /**
     * Sets aroundRadius
     *
     * @param \Algolia\AlgoliaSearch\Model\Search\AroundRadius|null $aroundRadius aroundRadius
     *
     * @return self
     */
    public function setAroundRadius($aroundRadius)
    {
        $this->container['aroundRadius'] = $aroundRadius;

        return $this;
    }

    /**
     * Gets insideBoundingBox
     *
     * @return float[]|null
     */
    public function getInsideBoundingBox()
    {
        return $this->container['insideBoundingBox'] ?? null;
    }

    /**
     * Sets insideBoundingBox
     *
     * @param float[]|null $insideBoundingBox search inside a rectangular area (in geo coordinates)
     *
     * @return self
     */
    public function setInsideBoundingBox($insideBoundingBox)
    {
        $this->container['insideBoundingBox'] = $insideBoundingBox;

        return $this;
    }

    /**
     * Gets insidePolygon
     *
     * @return float[]|null
     */
    public function getInsidePolygon()
    {
        return $this->container['insidePolygon'] ?? null;
    }

    /**
     * Sets insidePolygon
     *
     * @param float[]|null $insidePolygon search inside a polygon (in geo coordinates)
     *
     * @return self
     */
    public function setInsidePolygon($insidePolygon)
    {
        $this->container['insidePolygon'] = $insidePolygon;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param int $offset Offset
     *
     * @return bool
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param int $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param int $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }
}
