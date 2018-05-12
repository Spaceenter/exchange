package controller

// RequestInfo request information versioning controller and support mutiple protocals
type RequestInfo struct {
	Protocal   string //pbuf, json etc.
	APIVersion string //used for backward compatibility
}

//Query url query
type Query map[string]string

// 	public function ExceptionRouteHandler($method, $args)
// 	{
// 		$format = $this->protocol_format_;
// 		$this->protocol_format_ = Defined::Protocol_Json; // Exceptions always get JSON
// 		$unlock = true;

// 		$jails = \Bans\JailManager::GetInstance()->GetHook('FailedRequests')->GetJails();
// 		$failed = true;

// 		try {
// 			foreach ($jails as $jail) {
// 				$jail->CheckBan();
// 			}

// 			$response = $this->$method($args);
// 			$this->protocol_format_ = $format;

// 			$failed = false;
// 		}
// 		catch (InputError $e) {
// 			$response = Log::Exception($e, 'inputerror');
// 		}
// 		catch (WorkflowError $e) {
// 			$response = Log::Exception($e);
// 		}
// 		catch (RequestError $e) {
// 			$response = Log::Exception($e);
// 		}
// 		catch (HttpExceptionInterface $e) {
// 			$response = Log::Exception($e);
// 		}
// 		catch(\Bans\BanException $e) {
// 			// Ban stuff, put it in a middleware when we move to Slim V3
// 			$response = $e->GetResponse();

// 			$retryAfter = $e->GetRetryAfter();
// 			if (isset($retryAfter) && $retryAfter > 0) {
// 				$app = \Slim\Slim::getInstance();
// 				$app->response->headers->set( 'Retry-After',  $retryAfter);
// 			}
// 		}
// 		catch(\Throwable $e) {
// 			// This includes DatabaseErrors
// 			$unlock = false;
// 			// Return all unresolved exceptions if they spring up all the way up here
// 			$response = Log::Unhandled($e);

// 			// Set appropriate status code
// 			$app = \Slim\Slim::getInstance();
// 			$app->response->setStatus(Response::HTTP_SERVER_ERROR);
// 		}

// 		if ($unlock) {
// 			$this->UnlockObjects();
// 		}

// 		if ($failed) {
// 			foreach ($jails as $jail) {
// 				$jail->Add();
// 			}
// 		}

// 		return $response;
// 	}

// 	/**
// 	 * Process route arguments using FilterArray and ArrayObject representation
// 	 * @param  array $args      Input array with unspecified keys
// 	 * @param  array $allowed   List of allowed keys
// 	 * @return ArrayObject      Filtred array with allowed keys represented as ArrayObject propetries
// 	 */
// 	public static function Arguments($args, $required, $optional = null)
// 	{
// 		$filtered = Util::FilterArray($args, array_keys($required));

// 		if (count($filtered) !== count($required)) {
// 			throw RequestError::Response(
// 				Response::HTTP_BAD_REQUEST,
// 				tr("Error", "Invalid request parameters"));
// 		}

// 		if (!empty($optional)) {
// 			$append = Util::FilterArray($args, array_keys($optional));
// 			$filtered = array_merge($filtered, $append);
// 			$required = array_merge($required, $optional);
// 		}

// 		foreach ($required as $key => $type)
// 		{
// 			if ($type === 'integer') {
// 				$filtered[$key] = self::ProcessInteger($filtered, $key);
// 			} elseif ($type === 'string') {
// 				$filtered[$key] = self::ProcessString($filtered, $key);
// 			} elseif ($type === 'array') {
// 				$filtered[$key] = self::ProcessArray($filtered, $key);
// 			} elseif (is_callable($type)) {
// 				if (!isset($filtered[$key])) {
// 					$filtered[$key] = null;
// 					continue;
// 				}
// 				$filtered[$key] = $type($filtered[$key]);
// 			} elseif (is_array($type) && count(array_filter($type, 'is_callable')) === count($type)) {
// 				if (!isset($filtered[$key])) {
// 					$filtered[$key] = null;
// 					continue;
// 				}
// 				$applyFunc = function ($value, $func) { return $func($value); };
// 				$filtered[$key] = array_reduce($type, $applyFunc, isset($filtered[$key]) ? $filtered[$key] : null);
// 			} else {
// 				throw new \Exception("Unsupported type requested: ". var_export($type, true) ." for argument: {$key}");
// 			}
// 		}
// 		return new ArrayObject($filtered, ArrayObject::ARRAY_AS_PROPS);
// 	}

// 	public static function ProcessInteger($args, $key)
// 	{
// 		return isset($args[$key]) ? intval($args[$key]) : null;
// 	}

// 	public static function ProcessString($args, $key)
// 	{
// 		return isset($args[$key]) ? trim($args[$key]) : null;
// 	}

// 	public static function ProcessArray($args, $key)
// 	{
// 		return isset($args[$key]) ? $args[$key] : null;
// 	}

// 	protected function getUserID($user_name)
// 	{
// 		$user_id = User::get_user_id($user_name);

// 		if ($user_id > 0) {
// 			return $user_id;
// 		}

// 		if (strpos($user_name, '@') !== false) {

// 			$user_name = trim($user_name);

// 			// Strip domain if PM domain, people are stupid and confuse .ch and .com
// 			$split = explode('@', $user_name);
// 			$local = $split[0];
// 			$domain = $split[1];

// 			if (!in_array(strtolower($domain), AVAILABLE_DOMAINS)) {
// 				// Not a PM domain, forget it
// 				return 0;
// 			}

// 			$user_id = User::get_user_id($local);

// 			if ($user_id > 0) {

// 				$addresses = Address::get_addresses($user_id);
// 				$addresses = array_values(array_filter($addresses, function($element) { return $element->get_type() === Address::TYPE_ORIGINAL; }));

// 				if (empty($addresses)) {
// 					return 0;
// 				}

// 				$suggestion = $addresses[0]->get_email();

// 				throw RequestError::Response(Response::AUTH_ADDRESS_SUGGESTION, sprintf(tr("Error", "%s does not exist. Did you mean %s?"), $user_name, $suggestion));
// 			}
// 		}

// 		return 0;
// 	}

// 	protected function getAuthInfo($args)
// 	{
// 		if (isset($args['NewPassword']) || isset($args['Password'])) {
// 			$new_pwd = isset($args['NewPassword']) ? $args['NewPassword'] : $args['Password'];

// 			if (!isset($new_pwd) || !is_string($new_pwd)) {
// 				throw RequestError::Response(Response::AUTH_VALIDATE_PASSWORD_INPUT_INVALID, tr("Error", "Invalid input"));
// 			}

// 			if (strlen($new_pwd) == 0 || mb_strlen($new_pwd, 'UTF-8') > User::MAX_PASSWORD_LENGTH) {
// 				throw RequestError::Response(Response::AUTH_VALIDATE_PASSWORD_NEW_INVALID, tr("Error", "Invalid new password"));
// 			}

// 			return User::PasswordInfo($new_pwd);
// 		} else {

// 			if (!isset($args['Auth']) || !is_array($args['Auth'])) {
// 				throw RequestError::Response(Response::AUTH_VALIDATE_SRP_INPUT_INVALID, tr("Error", "Invalid input"));
// 			}

// 			$args = $args['Auth'];

// 			if (!isset($args['Version']) || !isset($args['ModulusID']) || !isset($args['Salt']) || !isset($args['Verifier'])
// 			  || !is_int($args['Version']) || !is_string($args['Salt']) || !is_string($args['Verifier'])) {
// 				throw RequestError::Response(Response::AUTH_VALIDATE_SRP_INPUT_INVALID, tr("Error", "Invalid input"));
// 			}

// 			if ($args['Version'] < User::MIN_AUTH_VERSION || $args['Version'] > User::AUTH_VERSION) {
// 				throw RequestError::Response(Response::APP_VERSION_BAD, tr("Error", "Application upgrade required"));
// 			}

// 			$srp_param_id = Util::decrypt_id($args['ModulusID']);
// 			if (!Util::is_id($srp_param_id)) {
// 				throw RequestError::Response(Response::AUTH_VALIDATE_SRP_INPUT_INVALID, tr("Error", "Invalid input"));
// 			}

// 			$salt = base64_decode($args['Salt']);
// 			$verifier = \ParagonIE\ConstantTime\Base64::decode($args['Verifier']);

// 			if (strlen($verifier) !== 2048 / 8 || strlen($salt) !== 10) {
// 				throw RequestError::Response(Response::AUTH_VALIDATE_SRP_INPUT_INVALID, tr("Error", "Invalid input"));
// 			}

// 			return array(
// 				'Version' => $args['Version'],
// 				'Salt' => $salt,
// 				'SRPVerifier' => $verifier,
// 				'SRPParamID' => $srp_param_id
// 			);
// 		}
// 	}

// 	protected function checkOrganizationKeys()
// 	{
// 		$public_key = Organization::get_current()->get_public_key();

// 		if (!isset($public_key)) {
// 			throw ResponseError::Response(Response::ORGANIZATION_KEYS_MISSING, tr("Error", "Multi-user mode must be enabled before performing this action"));
// 		}
// 	}

// 	public static function Callable(string $api_name, string $method_name, array $arg_names = array())
// 	{

// 		return function(...$args) use ($api_name, $method_name, $arg_names) {
// 			$request = RequestHandler::getInstance();

// 			$api = self::GetApi($api_name, $request->get_api_version());

// 			for($i=0; $i<count($arg_names); $i++) {
// 				if (!isset($args[$i])) {
// 					break;
// 				}

// 				$args[$arg_names[$i]] = $args[$i];
// 			}

// 			$body = $request->get_parsed_body();
// 			if (!empty($body) && is_array($body)) {
// 				$args = array_merge($args, $body);
// 			}

// 			// Only use query string params sent with GET, DELETE, and HEAD
// 			// This prevents the API call from working if clients put sensitive information in the query string
// 			$params = $request->get_request_params();
// 			$request_method = strtoupper($request->get_request_method());
// 			if (in_array($request_method, array('GET','DELETE','HEAD')) && !empty($params) && is_array($params)) {
// 				$args = array_merge($args, $params);
// 			}

// 			// FIXME, have a middleware handle this instead of echo here
// 			echo $api->ExceptionRouteHandler($method_name, $args);
// 		};
// 	}

// 	public static function GetApi($api_name, $api_version, $protocol = Defined::Protocol_Json)
// 	{
// 		$api_name = '\\App\\'.$api_name;
// 		$api_name_versioned = $api_name . 'V' . $api_version;
// 		if (class_exists($api_name_versioned)) {
// 			return new $api_name_versioned($protocol, $api_version);
// 		}
// 		return new $api_name($protocol, $api_version);
// 	}

// 	public static function Options()
// 	{
// 		// Nothing
// 	}
// }
// }
