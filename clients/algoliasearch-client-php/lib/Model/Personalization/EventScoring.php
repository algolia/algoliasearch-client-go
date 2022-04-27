<?php

namespace Algolia\AlgoliaSearch\Model\Personalization;

use Algolia\AlgoliaSearch\ObjectSerializer;

/**
 * EventScoring Class Doc Comment
 *
 * @category Class
 * @package Algolia\AlgoliaSearch
 */
class EventScoring extends \Algolia\AlgoliaSearch\Model\AbstractModel implements
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
        'score' => 'int',
        'eventName' => 'string',
        'eventType' => 'string',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'score' => null,
        'eventName' => null,
        'eventType' => null,
    ];

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
        'score' => 'setScore',
        'eventName' => 'setEventName',
        'eventType' => 'setEventType',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'score' => 'getScore',
        'eventName' => 'getEventName',
        'eventType' => 'getEventType',
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
        if (isset($data['score'])) {
            $this->container['score'] = $data['score'];
        }
        if (isset($data['eventName'])) {
            $this->container['eventName'] = $data['eventName'];
        }
        if (isset($data['eventType'])) {
            $this->container['eventType'] = $data['eventType'];
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

        if (
            !isset($this->container['score']) ||
            $this->container['score'] === null
        ) {
            $invalidProperties[] = "'score' can't be null";
        }
        if (
            !isset($this->container['eventName']) ||
            $this->container['eventName'] === null
        ) {
            $invalidProperties[] = "'eventName' can't be null";
        }
        if (
            !isset($this->container['eventType']) ||
            $this->container['eventType'] === null
        ) {
            $invalidProperties[] = "'eventType' can't be null";
        }
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
     * Gets score
     *
     * @return int
     */
    public function getScore()
    {
        return $this->container['score'] ?? null;
    }

    /**
     * Sets score
     *
     * @param int $score The score for the event.
     *
     * @return self
     */
    public function setScore($score)
    {
        $this->container['score'] = $score;

        return $this;
    }

    /**
     * Gets eventName
     *
     * @return string
     */
    public function getEventName()
    {
        return $this->container['eventName'] ?? null;
    }

    /**
     * Sets eventName
     *
     * @param string $eventName The name of the event.
     *
     * @return self
     */
    public function setEventName($eventName)
    {
        $this->container['eventName'] = $eventName;

        return $this;
    }

    /**
     * Gets eventType
     *
     * @return string
     */
    public function getEventType()
    {
        return $this->container['eventType'] ?? null;
    }

    /**
     * Sets eventType
     *
     * @param string $eventType The type of the event.
     *
     * @return self
     */
    public function setEventType($eventType)
    {
        $this->container['eventType'] = $eventType;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
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
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }
}
